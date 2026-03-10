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
		command := parts[0]
		args := parts[1:]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args, " "))
		default:
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
