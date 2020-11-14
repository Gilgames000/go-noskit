package packetsrv

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gilgames000/go-noskit/packets"

	"github.com/alecthomas/participle/lexer"
)

// NsTeST packet
type NsTeST struct {
	ServerNumber int              `json:"server_number" parser:"'NsTeST' @String"`
	Username     string           `json:"username" parser:"@String"`
	Unknown      int              `json:"unknown"       parser:"@String"`
	SessionID    int              `json:"session_id"    parser:"@String"`
	Endpoints    []ServerEndpoint `json:"endpoints"     parser:"@@*"`
}

// Name of the packet
func (p NsTeST) Name() string {
	return "NsTeST"
}

// Type of the packet
func (p NsTeST) Type() packets.PacketType {
	return packets.SERVER
}

// ServerEndpoint represents an element of the buy list
type ServerEndpoint struct {
	Address         string `json:"address"`
	Port            int    `json:"port"`
	ChannelColor    int    `json:"channel_color"`
	ChannelFullness int    `json:"channel_fullness"`
	ChannelNumber   int    `json:"channel_number"`
	ServerName      string `json:"server_name"`
}

func (se *ServerEndpoint) Parse(lex *lexer.PeekingLexer) error {
	token, err := lex.Next()
	if err != nil {
		return err
	}

	if token.EOF() {
		return errors.New("EOF reached")
	}

	split := strings.Split(token.Value, ":")

	se.Address = split[0]

	port, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	se.Port = port

	color, err := strconv.Atoi(split[2])
	if err != nil {
		return err
	}
	se.ChannelColor = color

	split = strings.Split(split[3], ".")

	fullness, err := strconv.Atoi(split[0])
	if err != nil {
		return err
	}
	se.ChannelFullness = fullness

	num, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	se.ChannelNumber = num

	se.ServerName = split[2]

	return nil
}
