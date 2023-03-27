package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandHelp(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("help")
	assertCommand(t, command, "help")
	assertArguments(t, arguments, []Argument{})
	// This just outputs to stdout. No specific test.
}
