package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
)

type CommandBin struct {
	resultFormatter resultformatter.ResultFormatter
}

func (commandBin *CommandBin) GetName() string {
	return "bin"
}

func (commandBin *CommandBin) GetSignatures() []Signature {
	return []Signature{
		// bin
		[]expreval.LexAnToken{}}
}

func (commandBin *CommandBin) Execute(arguments []Argument) error {
	commandBin.resultFormatter.SetOutputMode(resultformatter.OutputModeBinary)
	return nil
}

func (commandBin *CommandBin) GetUsage() (string, string) {
	return "bin", "Set binary output mode."
}

func NewCommandBin(resultFormatter resultformatter.ResultFormatter) Command {
	command := CommandBin{}
	command.resultFormatter = resultFormatter
	return &command
}
