package command

import (
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandExit(t *testing.T) {
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("exit")
	assertCommand(t, command, "exit")
	assertArguments(t, arguments, []Argument{})
}

func TestCommandExitWithTooManyArgs(t *testing.T) {
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(resultFormatter)
	command, _, err := commandParser.ParseCommand("exit 2 5")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
