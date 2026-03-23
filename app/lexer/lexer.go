package lexer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type Lexer struct {
	Reader bufio.Reader
	Tokens []Token
	cur    rune
	prev   rune
}

func (l *Lexer) NextToken() (*Token, error) {
	l.skipWhiteSpace()
	if l.cur == 0 {
		return nil, io.EOF
	}

	switch l.cur {
	default:
		word := l.readWord()
		if word == "" {
			return nil, io.EOF
		}
		return NewToken(Word, word), nil
	}
}

func (l *Lexer) advance() {
	r, _, err := l.Reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.prev = l.cur
			l.cur = 0
			return
		}
		panic(fmt.Sprintf("error reading rune: %v", err))
	}

	l.prev = l.cur
	l.cur = r
}

func (l *Lexer) readWord() string {
	if l.cur == '\'' {
		l.advance()
		word := ""
		for l.cur != 0 && l.cur != '\'' {
			word += string(l.cur)
			l.advance()
		}
		if l.cur == '\'' {
			l.advance()
		}
		return word
	}
	word := ""
	for l.cur != 0 && !unicode.IsSpace(l.cur) {
		word += string(l.cur)
		l.advance()
	}
	return word
}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.cur) && l.cur != 0 {
		l.advance()
	}
}

func NewLexer(reader bufio.Reader) *Lexer {
	r, _, err := reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return &Lexer{Reader: reader, Tokens: []Token{}, cur: 0}
		}
		panic(fmt.Sprintf("error reading first rune: %v", err))
	}
	return &Lexer{Reader: reader, Tokens: []Token{}, cur: r}
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
