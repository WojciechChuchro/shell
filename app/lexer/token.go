package lexer

import "fmt"

type TokenType int

const (
	Word TokenType = iota
)

func (tt TokenType) String() string {
	switch tt {
	case Word:
		return "Word"
	default:
		return fmt.Sprintf("Unknown(%d)", int(tt))
	}
}

type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s", t.Type.String(), t.Value)
}

func NewToken(tokenType TokenType, value string) *Token {
	return &Token{
		Type:  tokenType,
		Value: value,
	}
}
