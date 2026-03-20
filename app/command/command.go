package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command interface {
	Execute() error
}

type SimpleCommand struct {
	Name string
	Args []string
}

func (c *SimpleCommand) Execute() error {
	switch c.Name {
	case "pwd":
		return pwd()
	case "cd":
		return cd(c.Args)
	case "echo":
		return echo(c.Args)
	case "type":
		return typeCmd(c.Args)
	}
	return runExternal(c.Name, c.Args)
}

func pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("pwd: error getting current directory")
	}
	fmt.Println(dir)
	return nil
}

func cd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: missing argument")
	}

	path := args[0]
	if path == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cd: could not resolve home directory")
		}
		path = home
	}

	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", args[0])
	}
	return nil
}

func echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func typeCmd(args []string) error {
	builtins := map[string]struct{}{
		"exit": {},
		"echo": {},
		"type": {},
		"pwd":  {},
		"cd":   {},
	}

	for _, arg := range args {
		if _, ok := builtins[arg]; ok {
			fmt.Printf("%s is a shell builtin\n", arg)
			continue
		}

		path, err := exec.LookPath(arg)
		if err != nil {
			fmt.Printf("%s: not found\n", arg)
			continue
		}
		fmt.Printf("%s is %s\n", arg, path)
	}
	return nil
}

func runExternal(name string, args []string) error {
	if _, err := exec.LookPath(name); err != nil {
		fmt.Printf("%s: command not found\n", name)
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}