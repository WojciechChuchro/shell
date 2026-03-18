package parser

import (
	"bufio"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

type Parser struct {
	reader *bufio.Reader
}

func (p *Parser) Parse() *command.Command {
	return nil
}
