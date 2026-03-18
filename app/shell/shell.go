package shell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/lexer"
)

type Shell struct {
}

func (s *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)
	lexer := lexer.NewLexer()
	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command")
			continue
		}
		lexer.Parse(input)

	}
}

func NewShell() *Shell {
	return &Shell{}
}
