package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	Reader bufio.Reader
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
	var word strings.Builder

	for l.cur != 0 && !unicode.IsSpace(l.cur) {
		if l.cur == '\'' {
			l.advance()
			for l.cur != 0 && l.cur != '\'' {
				word.WriteRune(l.cur)
				l.advance()
			}
			if l.cur == '\'' {
				l.advance()
			}
			continue
		}

		if l.cur == '"' {
			l.advance()
			for l.cur != 0 && l.cur != '"' {
				word.WriteRune(l.cur)
				l.advance()
			}
			if l.cur == '"' {
				l.advance()
			}
			continue
		}

		word.WriteRune(l.cur)
		l.advance()
	}
	return word.String()
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
			return &Lexer{Reader: reader, cur: 0}
		}
		panic(fmt.Sprintf("error reading first rune: %v", err))
	}
	return &Lexer{Reader: reader, cur: r}
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
