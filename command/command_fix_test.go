package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandFixNoPrecision(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("fix")
	assertCommand(t, command, "fix")
	assertArguments(t, arguments, []Argument{})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeFixed, 2)
}

func TestCommandFixWithPrecision(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("fix 4")
	assertCommand(t, command, "fix")
	assertArguments(t, arguments, []Argument{{expreval.TokenNumber, "", 4.0}})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeFixed, 4)
}

func TestCommandFixWithTooManyArgs(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, _, err := commandParser.ParseCommand("fix 2 5")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
