package command

import (
	"alanmitic/gocalc/expreval"
	"fmt"
)

type CommandVars struct {
	evaluator *expreval.Evaluator
}

func (commandVar *CommandVars) GetName() string {
	return "vars"
}

func (commandVar *CommandVars) GetSignatures() []Signature {
	return []Signature{
		// vars
		[]expreval.LexAnToken{}}
}

func (commandVar *CommandVars) Execute(arguments []Argument) error {
	variables := commandVar.evaluator.VariableStore
	if len(variables) > 0 {
		fmt.Println("Variables:")
		for variableName, variable := range variables {
			fmt.Println(variableName, " => ", variable)
		}
	} else {
		fmt.Println("No variables defined!")
	}

	return nil
}

func (commandVar *CommandVars) GetUsage() (string, string) {
	return "vars", "List defined variables."
}

func NewCommandVars(evaluator *expreval.Evaluator) Command {
	command := CommandVars{}
	command.evaluator = evaluator
	return &command
}
