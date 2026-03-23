package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/lexer"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
)

type Shell struct {
}

func (s *Shell) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		line = strings.TrimSpace(line)
		if line != "" {
			if line == "exit 0" || line == "exit" {
				return
			}

			lineReader := bufio.NewReader(strings.NewReader(line))
			lexer := lexer.NewLexer(*lineReader)
			parser := parser.NewParser(lexer)
			cmd := parser.Parse()
			if cmd != nil {
				if err := cmd.Execute(); err != nil {
					fmt.Println(err)
				}
			}
		}

		if err == io.EOF {
			return
		}
	}
}

func NewShell() *Shell {
	return &Shell{}
}
