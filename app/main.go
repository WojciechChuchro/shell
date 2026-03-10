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
		command := strings.TrimSpace(line)

		switch command {
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
