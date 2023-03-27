package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
)

type CommandHex struct {
	resultFormatter resultformatter.ResultFormatter
}

func (commandHex *CommandHex) GetName() string {
	return "hex"
}

func (commandHex *CommandHex) GetSignatures() []Signature {
	return []Signature{
		// hex
		[]expreval.LexAnToken{}}
}

func (commandHex *CommandHex) Execute(arguments []Argument) error {
	commandHex.resultFormatter.SetOutputMode(resultformatter.OutputModeHexadecimal)
	return nil
}

func (commandHex *CommandHex) GetUsage() (string, string) {
	return "hex", "Set hexadecimal output mode."
}

func NewCommandHex(resultFormatter resultformatter.ResultFormatter) Command {
	command := CommandHex{}
	command.resultFormatter = resultFormatter
	return &command
}
