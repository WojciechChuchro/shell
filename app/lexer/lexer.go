package lexer

import "strings"

type Lexer struct {
	Tokens []Token
	start  int
}

func (l *Lexer) Parse(input string) ([]Token, error) {
	input = strings.TrimSpace(input)

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			l.Append(input, i)
		} else if i == len(input)-1 {
			l.Append(input, i+1)
		}
	}

	return l.Tokens, nil
}

func NewLexer() *Lexer {
	return &Lexer{}
}

func (l *Lexer) Append(input string, i int) {
	token := NewToken(Word, input[l.start:i])
	l.Tokens = append(l.Tokens, *token)
	l.start = i + 1
}
