package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandOct(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("oct")
	assertCommand(t, command, "oct")
	assertArguments(t, arguments, []Argument{})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeOctal, -1)
}

func TestCommandOctWithTooManyArgs(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, _, err := commandParser.ParseCommand("oct 2")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
