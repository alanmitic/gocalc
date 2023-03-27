package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
)

const (
	DefaultFixPrecision = 2
)

type CommandFix struct {
	resultFormatter resultformatter.ResultFormatter
}

func (commandFix *CommandFix) GetName() string {
	return "fix"
}

func (commandFix *CommandFix) GetSignatures() []Signature {
	return []Signature{
		// fix
		[]expreval.LexAnToken{},
		// fix <n>
		[]expreval.LexAnToken{expreval.TokenNumber}}
}

func (commandFix *CommandFix) Execute(arguments []Argument) error {
	precision := DefaultFixPrecision
	if len(arguments) == 1 {
		precision = int(arguments[0].numericValue)
	}
	commandFix.resultFormatter.SetOutputMode(resultformatter.OutputModeFixed)
	commandFix.resultFormatter.SetPrecision(precision)
	return nil
}

func (commandFix *CommandFix) GetUsage() (string, string) {
	return "fix <precision>", "Set fix point output mode with optional precision."
}

func NewCommandFix(resultFormatter resultformatter.ResultFormatter) Command {
	command := CommandFix{}
	command.resultFormatter = resultFormatter
	return &command
}
