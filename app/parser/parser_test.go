package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/command"
	"github.com/codecrafters-io/shell-starter-go/app/lexer"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected command.SimpleCommand
	}{
		{
			name: "single command",
			input: "ls",
			expected: command.SimpleCommand{Name: "ls", Args: []string{}},
		},
		{
			name: "command with flag",
			input: "ls -la",
			expected: command.SimpleCommand{Name: "ls", Args: []string{"-la"}},
		},
		{
			name: "command with quotes",
			input: "ls 'hello world'",
			expected: command.SimpleCommand{Name: "ls", Args: []string{"hello world"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(test.input))
			lexer := lexer.NewLexer(*reader)
			parser := NewParser(lexer)
			cmd := parser.Parse()
			if cmd == nil {
				t.Fatal("expected command, got nil")
			}

			simpleCmd, ok := cmd.(*command.SimpleCommand)
			if !ok {
				t.Fatalf("expected *command.SimpleCommand, got %T", cmd)
			}

			if simpleCmd.Name != test.expected.Name {
				t.Fatalf("expected command name %q, got %q", test.expected.Name, simpleCmd.Name)
			}
			if len(simpleCmd.Args) != len(test.expected.Args) {
				t.Fatalf("expected %d args, got %d", len(test.expected.Args), len(simpleCmd.Args))
			}
			for i := range simpleCmd.Args {
				if simpleCmd.Args[i] != test.expected.Args[i] {
					t.Fatalf("expected arg %q, got %q", test.expected.Args[i], simpleCmd.Args[i])
				}
			}
		})
	}
}