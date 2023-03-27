package command

import (
	"alanmitic/gocalc/expreval"
	"fmt"
	"strconv"
)

type CommandHelp struct {
	commands map[string]Command
}

func (commandHelp *CommandHelp) GetName() string {
	return "help"
}

func (commandHelp *CommandHelp) GetSignatures() []Signature {
	return []Signature{
		// help
		[]expreval.LexAnToken{}}
}

func (commandHelp *CommandHelp) Execute(arguments []Argument) error {
	commands := commandHelp.commands

	longestUsageSyntaxLength := findLongestUsageSyntaxLength(commands)

	fmt.Println("Commands:")
	for _, command := range commands {
		usageSyntax, usageDescription := command.GetUsage()

		paddedUsageSyntax := fmt.Sprintf("% -"+strconv.Itoa(longestUsageSyntaxLength)+"s", usageSyntax)
		fmt.Println(paddedUsageSyntax, ":", usageDescription)
	}

	return nil
}

func findLongestUsageSyntaxLength(commands map[string]Command) int {
	longestUsageLen := 0

	for _, command := range commands {
		usageSyntax, _ := command.GetUsage()
		usageLen := len(usageSyntax)
		if usageLen > longestUsageLen {
			longestUsageLen = usageLen
		}

	}

	return longestUsageLen
}

func (commandHelp *CommandHelp) GetUsage() (string, string) {
	return "help", "Show help"
}

func NewCommandHelp(commands map[string]Command) Command {
	command := CommandHelp{}
	command.commands = commands
	return &command
}
