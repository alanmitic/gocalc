package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
)

const (
	DefaultSciPrecision = 2
)

type CommandSci struct {
	resultFormatter resultformatter.ResultFormatter
}

func (commandSci *CommandSci) GetName() string {
	return "sci"
}

func (commandSci *CommandSci) GetSignatures() []Signature {
	return []Signature{
		// sci
		[]expreval.LexAnToken{},
		// sci <n>
		[]expreval.LexAnToken{expreval.TokenNumber}}
}

func (commandSci *CommandSci) Execute(arguments []Argument) error {
	precision := DefaultSciPrecision
	if len(arguments) == 1 {
		precision = int(arguments[0].numericValue)
	}
	commandSci.resultFormatter.SetOutputMode(resultformatter.OutputModeScientific)
	commandSci.resultFormatter.SetPrecision(precision)
	return nil
}

func (commandSci *CommandSci) GetUsage() (string, string) {
	return "sci <precision>", "Set scientific output mode with optional precision."
}

func NewCommandSci(resultFormatter resultformatter.ResultFormatter) Command {
	command := CommandSci{}
	command.resultFormatter = resultFormatter
	return &command
}
