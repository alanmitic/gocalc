package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func TestCommandVars(t *testing.T) {
	evaluator := expreval.NewEvaluator()
	resultFormatter := resultformatter.NewResultFormatter()
	commandParser := NewCommandParser(evaluator, resultFormatter)
	command, arguments, _ := commandParser.ParseCommand("vars")
	assertCommand(t, command, "vars")
	assertArguments(t, arguments, []Argument{})
	// This just outputs to stdout. No specific test.
}
