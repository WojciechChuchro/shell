package lexer

import "fmt"

type TokenType int

const (
	Word TokenType = iota
)

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("%x %s", t.Type, t.Value)
}

func NewToken(tokenType TokenType, value string) *Token {
	return &Token{
		Type:  tokenType,
		Value: value,
	}
}
