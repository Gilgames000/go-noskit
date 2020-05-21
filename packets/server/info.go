package packetsrv

import (
	"strings"

	"github.com/gilgames000/go-noskit/errors"
	"github.com/gilgames000/go-noskit/packets"

	"github.com/alecthomas/participle/lexer"
)

// Info packet
type Info struct {
	Message string `json:"message"`
}

// Name of the packet
func (p Info) Name() string {
	return "info"
}

// Type of the packet
func (p Info) Type() packets.PacketType {
	return packets.SERVER
}

func (p *Info) Parse(lex *lexer.PeekingLexer) error {
	token, err := lex.Next()
	if err != nil {
		return err
	}

	if token.Value != "info" {
		return errors.Errorf("syntax error 'info' expected, but got %s", token.Value)
	}

	token, err = lex.Next()
	if err != nil {
		return err
	}
	msg := strings.Builder{}
	msg.WriteString(token.Value)

	for {
		token, err = lex.Next()
		if err != nil {
			return err
		}

		if token.EOF() {
			break
		}

		msg.WriteByte(' ')
		msg.WriteString(token.Value)
	}
	p.Message = msg.String()

	return nil
}
