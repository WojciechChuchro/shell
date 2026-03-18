package lexer

import (
	"testing"
)

func AssertToken(tokenType TokenType, value string, token Token, t *testing.T) {
	if token.Type != tokenType {
		t.Errorf("Expected token type to be Word, got %d", token.Type)
	}
	if token.Value != value {
		t.Errorf("Expected token value to be ls, got %s", token.Value)
	}

}

func TestLexer(t *testing.T) {
	input := "ls"

	s := NewLexer()
	tokens, err := s.Parse(input)

	if err != nil {
		t.Errorf("Error parsing input: %s", err)
	}

	AssertToken(Word, "ls", tokens[0], t)
}
