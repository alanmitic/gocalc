package command

import (
	"alanmitic/gocalc/expreval"
	"os"
)

type CommandExit struct {
}

func (commandExit *CommandExit) GetName() string {
	return "exit"
}

func (commandExit *CommandExit) GetSignatures() []Signature {
	return []Signature{
		// exit
		[]expreval.LexAnToken{}}
}

func (commandExit *CommandExit) Execute(arguments []Argument) error {
	os.Exit(0)
	return nil
}

func NewCommandExit() Command {
	command := CommandExit{}
	return &command
}
