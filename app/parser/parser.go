package parser

import (
	"bufio"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/lexer"
)

type Parser struct {
	reader *bufio.Reader
}

func (p *Parser) Parse() *command.Command {
	lexer := lexer.NewLexer(p.reader)
}
