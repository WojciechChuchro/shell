package parser

import (
	"io"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
	tokens []lexer.Token
}

func (p *Parser) Parse() command.Command {
	p.tokens = p.tokens[:0]
	for {
		token, err := p.lexer.NextToken()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		p.tokens = append(p.tokens, *token)
	}

	if len(p.tokens) == 0 {
		return nil
	}

	args := make([]string, 0, len(p.tokens)-1)
	for i := 1; i < len(p.tokens); i++ {
		args = append(args, p.tokens[i].Value)
	}

	return &command.SimpleCommand{
		Name: p.tokens[0].Value,
		Args: args,
	}
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer}
}