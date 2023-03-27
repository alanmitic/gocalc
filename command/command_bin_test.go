package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandBin(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("bin")
	assertCommand(t, command, "bin")
	assertArguments(t, arguments, []Argument{})

	command.Execute(arguments)
	assertOutputModeAndPrecision(t, resultFormatter, resultformatter.OutputModeBinary, -1)
}

func TestCommandBinWithTooManyArgs(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, _, err := commandParser.ParseCommand("bin 2")
	assertNilCommandAndError(t, command, err, ErrTooManyArgs)
}
