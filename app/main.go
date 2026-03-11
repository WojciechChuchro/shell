package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	builtins := map[string]struct{}{
		"exit": {},
		"echo": {},
		"type": {},
	}
	for {
		fmt.Print("$ ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		parts := strings.Split(strings.TrimSpace(line), " ")
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
					continue
				}

				if path, err := exec.LookPath(arg); err == nil {
					fmt.Printf("%s is %s\n", arg, path)
				} else {
					fmt.Printf("%s: not found\n", arg)
				}
			}
		default:
			// External command: reuse LookPath logic, but keep argv[0] as the command name
			if _, err := exec.LookPath(command); err == nil {
				cmd := exec.Command(command, args...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				_ = cmd.Run()
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}

	}
}
