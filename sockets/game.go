package sockets

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/gilgames000/go-noskit/errors"
	"github.com/gilgames000/go-noskit/gamestate"
	"github.com/gilgames000/go-noskit/packets"
	"github.com/gilgames000/go-noskit/packets/parser"

	"github.com/gilgames000/go-noscrypto/pkg/noscryptoclt"
)

var _ gamestate.GameSocket = &GameSocket{}

type GameSocket struct {
	conn            net.Conn
	done            chan struct{}
	packetsToBeSent chan string

	connectedMtx sync.RWMutex
	isConnected  bool

	sessionNumber int
	packetNumber  uint16

	parser        parser.NosPacketParser
	packetsPubSub *NosPacketPubSub
}

func NewGameSocket(parser parser.NosPacketParser) *GameSocket {
	return &GameSocket{
		parser:        parser,
		packetsPubSub: NewNosPacketPubSub(),
	}
}

func (gs *GameSocket) Connect(address string, sessionNumber int) error {
	if gs.isConnected {
		return errors.New("game socket already connected")
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return errors.Wrap(err, "failed to connect to the game server")
	}

	gs.packetNumber = uint16(rand.Intn(65535))
	sessionPacket := noscryptoclt.EncryptSessionPacket(strconv.Itoa(sessionNumber))
	conn.Write([]byte(sessionPacket))
	gs.packetNumber++

	gs.conn = conn
	gs.done = make(chan struct{})
	gs.packetsToBeSent = make(chan string)
	go gs.sender()
	go gs.receiver()
	gs.isConnected = true

	return nil
}

func (gs *GameSocket) Disconnect() error {
	if !gs.isConnected {
		return errors.New("can't disconnect game socket because it is not connected")
	}

	gs.connectedMtx.Lock()
	gs.isConnected = false
	gs.connectedMtx.Unlock()
	close(gs.done)

	return gs.conn.Close()
}

func (gs *GameSocket) IsConnected() bool {
	gs.connectedMtx.RLock()
	connected := gs.isConnected
	gs.connectedMtx.RUnlock()

	return connected
}

func (gs *GameSocket) Send(packet ...packets.NosPacketStringer) error {
	if !gs.IsConnected() {
		return errors.New("can't send packet because game socket is not connected")
	}

	for _, p := range packet {
		gs.packetsToBeSent <- p.String()
	}

	return nil
}

func (gs *GameSocket) SendRaw(packet ...string) error {
	if !gs.IsConnected() {
		return errors.New("can't send packet because game socket is not connected")
	}

	for _, p := range packet {
		gs.packetsToBeSent <- p
	}

	return nil
}

func (gs *GameSocket) NewListener(packetNames ...string) <-chan packets.NosPacket {
	return gs.packetsPubSub.Subscribe(packetNames...)
}

func (gs *GameSocket) CloseListener(listener <-chan packets.NosPacket) {
	gs.packetsPubSub.Unsubscribe(listener)
}

func (gs *GameSocket) sender() {
	for {
		select {
		case p := <-gs.packetsToBeSent:
			packet := fmt.Sprintf("%d %s", gs.packetNumber, p)
			encryptedPacket := noscryptoclt.EncryptGamePacket(packet, gs.sessionNumber)
			gs.conn.Write([]byte(encryptedPacket))
			gs.packetNumber++
		case <-gs.done:
			return
		}
	}
}

func (gs *GameSocket) receiver() {
	var packet string

	scanner := bufio.NewScanner(gs.conn)
	scanner.Split(scanFF)
	for scanner.Scan() {
		packet = noscryptoclt.DecryptGamePacket(scanner.Text())
		packet = strings.TrimRight(packet, "\r\n")

		parsedPacket, err := gs.parser.ParseServerPacket(packet)
		if err != nil {
			continue
		}
		gs.packetsPubSub.Publish(parsedPacket)
	}
}

func scanFF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\xFF'); i >= 0 {
		// We have a full 0xFF-terminated line.
		return i + 1, data[0 : i+1], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
