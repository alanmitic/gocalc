package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
)

type CommandOct struct {
	resultFormatter resultformatter.ResultFormatter
}

func (commandOct *CommandOct) GetName() string {
	return "oct"
}

func (commandOct *CommandOct) GetSignatures() []Signature {
	return []Signature{
		// oct
		[]expreval.LexAnToken{}}
}

func (commandOct *CommandOct) Execute(arguments []Argument) error {
	commandOct.resultFormatter.SetOutputMode(resultformatter.OutputModeOctal)
	return nil
}

func (commandOct *CommandOct) GetUsage() (string, string) {
	return "oct", "Set octal output mode."
}

func NewCommandOct(resultFormatter resultformatter.ResultFormatter) Command {
	command := CommandOct{}
	command.resultFormatter = resultFormatter
	return &command
}
