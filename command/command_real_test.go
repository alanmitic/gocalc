package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandRealNoPrecision(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("real")
	assertCommand(t, command, "real")
	assertArguments(t, arguments, []Argument{})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeReal, -1)
}

func TestCommandRealWithPrecision(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("real 4")
	assertCommand(t, command, "real")
	assertArguments(t, arguments, []Argument{{expreval.TokenNumber, "", 4.0}})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeReal, 4)
}

func TestCommandRealWithTooManyArgs(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, _, err := commandParser.ParseCommand("real 2 5")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
