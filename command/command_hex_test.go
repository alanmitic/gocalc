package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandHex(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("hex")
	assertCommand(t, command, "hex")
	assertArguments(t, arguments, []Argument{})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeHexadecimal, -1)
}

func TestCommandHexWithTooManyArgs(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, _, err := commandParser.ParseCommand("hex 2")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
