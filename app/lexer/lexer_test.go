package lexer

import (
	"bufio"
	"strings"
	"testing"
)

func assertToken(expected, token Token, t *testing.T) {
	t.Helper()

	if token.Type != expected.Type {
		t.Errorf("expected token type %v, got %v", expected.Type, token.Type)
	}

	if token.Value != expected.Value {
		t.Errorf("expected token value %q, got %q", expected.Value, token.Value)
	}
}

func TestLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:  "single command",
			input: "ls",
			expected: []Token{
				{Type: Word, Value: "ls"},
			},
		},
		{
			name:  "command with flag",
			input: "ls -la",
			expected: []Token{
				{Type: Word, Value: "ls"},
				{Type: Word, Value: "-la"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(test.input))
			lexer := NewLexer(*reader)
			makeTokens := func() []*Token {
				tokens := []*Token{}
				for {
					token, err := lexer.NextToken()
					if err != nil {
						break
					}
					tokens = append(tokens, token)
				}
				return tokens
			}
			tokens := makeTokens()
			if len(tokens) != len(test.expected) {
				t.Fatalf("expected %d tokens, got %d", len(test.expected), len(tokens))
			}
			for i := range tokens {
				assertToken(test.expected[i], *tokens[i], t)
			}
		})
	}
}
