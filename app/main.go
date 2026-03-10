package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		parts := strings.Split(strings.TrimSpace(line), " ")
		builtins := make(map[string]struct{})
		builtins["exit"] = struct{}{}
		builtins["echo"] = struct{}{}
		builtins["type"] = struct{}{}
		command := parts[0]
		args := parts[1:]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "type":
			for _, arg := range args {
				if _, ok := builtins[arg]; ok {
					fmt.Printf("%s is a shell builtin\n", arg)
				} else {
					fmt.Printf("%s: not found\n", arg)
				}
			}
		default:
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
