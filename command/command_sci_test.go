package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandSciNoPrecision(t *testing.T) {
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("sci")
	assertCommand(t, command, "sci")
	assertArguments(t, arguments, []Argument{})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeScientific, 2)
}

func TestCommandSciWithPrecision(t *testing.T) {
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("sci 4")
	assertCommand(t, command, "sci")
	assertArguments(t, arguments, []Argument{{expreval.TokenNumber, "", 4.0}})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeScientific, 4)
}

func TestCommandSciWithTooManyArgs(t *testing.T) {
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(resultFormatter)
	command, _, err := commandParser.ParseCommand("sci 2 5")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
