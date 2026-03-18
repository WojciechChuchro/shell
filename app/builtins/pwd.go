package builtins

import "github.com/codecrafters-io/shell-starter-go/app/command"

type Pwd struct {
}

func (p *Pwd) Execute() error {
	return nil
}

func NewPwd() command.Command {
	return &Pwd{}
}
