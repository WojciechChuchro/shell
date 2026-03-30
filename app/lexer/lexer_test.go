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
		{
			name:  "command with quotes",
			input: "ls 'hello   world'",
			expected: []Token{
				{Type: Word, Value: "ls"},
				{Type: Word, Value: "hello   world"},
			},
		},
		{
			name:  "concatenated quoted and unquoted fragments",
			input: "echo 'test     shell' 'example''script' world''hello",
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: "test     shell"},
				{Type: Word, Value: "examplescript"},
				{Type: Word, Value: "worldhello"},
			},
		},
		{
			name:  "double quotes preserve spaces",
			input: `echo "hello    world"`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: "hello    world"},
			},
		},
		{
			name:  "adjacent double quoted strings concatenate",
			input: `echo "hello""world"`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: "helloworld"},
			},
		},
		{
			name:  "separate double quoted arguments",
			input: `echo "hello" "world"`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: "hello"},
				{Type: Word, Value: "world"},
			},
		},
		{
			name:  "single quote inside double quotes is literal",
			input: `echo "shell's test"`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: "shell's test"},
			},
		},
		{
			name:  "escaped spaces",
			input: "echo three\\ \\ \\ spaces",
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: "three   spaces"},
			},
		},
		{
			name:  "escaped single quotes",
			input: `echo three'\\ \\ \\' spaces`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: `three\\ \\ \\`},
				{Type: Word, Value: "spaces"},
			},
		},
		{
			name:  "escaped double quotes",
			input: `echo "script'shell'\\'example"`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: `script'shell'\'example`},
			},
		},
		{
			name:  "backslash before non-special char in double quotes is preserved",
			input: `echo "foo\bar"`,
			expected: []Token{
				{Type: Word, Value: "echo"},
				{Type: Word, Value: `foo\bar`},
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
