package parser

import (
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
)

// NosLexer returns a lexer instance able to tokenize nostale packets
func NosLexer() lexer.Definition {
	return lexer.Must(regex.New(`
		Whitespace = \s+
		String     = [^\s]+
	`))
}
