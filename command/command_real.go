package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
)

const (
	DefaultRealPrecision = -1
)

type CommandReal struct {
	resultFormatter resultformatter.ResultFormatter
}

func (commandReal *CommandReal) GetName() string {
	return "real"
}

func (commandReal *CommandReal) GetSignatures() []Signature {
	return []Signature{
		// real
		[]expreval.LexAnToken{},
		// real <n>
		[]expreval.LexAnToken{expreval.TokenNumber}}
}

func (commandReal *CommandReal) Execute(arguments []Argument) error {
	precision := DefaultRealPrecision
	if len(arguments) == 1 {
		precision = int(arguments[0].numericValue)
	}
	commandReal.resultFormatter.SetOutputMode(resultformatter.OutputModeReal)
	commandReal.resultFormatter.SetPrecision(precision)
	return nil
}

func NewCommandReal(resultFormatter resultformatter.ResultFormatter) Command {
	command := CommandReal{}
	command.resultFormatter = resultFormatter
	return &command
}
