package lexer

import (
	"bufio"
	"fmt"
	"io"
)

type Lexer struct {
	Reader bufio.Reader
	Tokens []Token
	cur    rune
	prev   rune
}

func (l *Lexer) NextToken() (Token, error) {
	l.skipWhiteSpace()

	switch l.cur {
	default:
		l.readWord()
		return Token{}, io.EOF
	}

}

func (l *Lexer) NextRune() {
	r, _, err := l.Reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("error reading rune: %v", err))
	}

	l.prev = l.cur
	l.cur = r
}

func (l *Lexer) readWord() {

}

func (l *Lexer) skipWhiteSpace() {
	for l.cur != ' ' {
		l.NextRune()
	}
}

func NewLexer(reader bufio.Reader) *Lexer {
	r, _, err := reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("error reading first rune: %v", err))
	}
	return &Lexer{Reader: reader, Tokens: []Token{}, cur: r}
}
