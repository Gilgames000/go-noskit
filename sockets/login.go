package sockets

import (
	"bufio"
	"net"
	"strings"
	"sync"

	"github.com/gilgames000/go-noskit/errors"
	"github.com/gilgames000/go-noskit/gamestate"
	"github.com/gilgames000/go-noskit/packets"
	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/gilgames000/go-noscrypto/pkg/noscryptoclt"
)

var _ gamestate.LoginSocket = &LoginSocket{}

type LoginSocket struct {
	conn            net.Conn
	done            chan struct{}
	packetsToBeSent chan string

	connectedMtx sync.RWMutex
	isConnected  bool

	parser        parser.NosPacketParser
	packetsPubSub *NosPacketPubSub
}

func NewLoginSocket(parser parser.NosPacketParser) *LoginSocket {
	return &LoginSocket{
		parser:        parser,
		packetsPubSub: NewNosPacketPubSub(),
	}
}

func (ls *LoginSocket) Connect(address string) error {
	if ls.isConnected {
		return errors.New("login socket already connected")
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return errors.Wrap(err, "failed to connect to the login server")
	}

	ls.conn = conn
	ls.done = make(chan struct{})
	ls.packetsToBeSent = make(chan string)
	go ls.sender()
	go ls.receiver()
	ls.isConnected = true

	return nil
}

func (ls *LoginSocket) Disconnect() error {
	if !ls.isConnected {
		return errors.New("can't disconnect login socket because it is not connected")
	}

	ls.connectedMtx.Lock()
	ls.isConnected = false
	ls.connectedMtx.Unlock()
	close(ls.done)

	return ls.conn.Close()
}

func (ls *LoginSocket) IsConnected() bool {
	ls.connectedMtx.RLock()
	connected := ls.isConnected
	ls.connectedMtx.RUnlock()

	return connected
}

func (ls *LoginSocket) Send(packet ...packets.NosPacketStringer) error {
	if !ls.IsConnected() {
		return errors.New("can't send packet because login socket is not connected")
	}

	for _, p := range packet {
		ls.packetsToBeSent <- p.String()
	}

	return nil
}

func (ls *LoginSocket) SendRaw(packet ...string) error {
	if !ls.IsConnected() {
		return errors.New("can't send packet because login socket is not connected")
	}

	for _, p := range packet {
		ls.packetsToBeSent <- p
	}

	return nil
}

func (ls *LoginSocket) NewListener(packetNames ...string) <-chan packets.NosPacket {
	return ls.packetsPubSub.Subscribe(packetNames...)
}

func (ls *LoginSocket) CloseListener(listener <-chan packets.NosPacket) {
	ls.packetsPubSub.Unsubscribe(listener)
}

func (ls *LoginSocket) sender() {
	for {
		select {
		case p := <-ls.packetsToBeSent:
			encryptedPacket := noscryptoclt.EncryptLoginPacket(p)
			ls.conn.Write([]byte(encryptedPacket))
		case <-ls.done:
			return
		}
	}
}

func (ls *LoginSocket) receiver() {
	var packet string

	scanner := bufio.NewScanner(ls.conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		packet = noscryptoclt.DecryptLoginPacket(scanner.Text())
		packet = strings.TrimRight(packet, "\r\n")

		parsedPacket, err := ls.parser.ParseServerPacket(packet)
		if err != nil {
			continue
		}
		ls.packetsPubSub.Publish(parsedPacket)
	}
}
