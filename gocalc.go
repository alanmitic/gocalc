package main

import (
	"alanmitic/gocalc/command"
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"bufio"
	"fmt"
	"os"
)

func main() {
	evaluator := expreval.NewEvaluator()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("gocalc version 0.0.1\n\nType 'exit' and ENTER to quit.")

	resultformatter := resultformatter.NewResultFormatter()
	commandParser := command.NewCommandParser(resultformatter)

	for {
		fmt.Print("gocalc >> ")
		exprOrCmd, _ := reader.ReadString('\n')

		cmd, arguments, err := commandParser.ParseCommand(exprOrCmd)
		if err != nil {
			fmt.Println("COMMAND ERROR:", err)
			// TODO: Output command usage.
			continue
		}

		if cmd != nil {
			cmd.Execute(arguments)
		} else {
			result, err := evaluator.Evaluate(exprOrCmd)
			if err != nil {
				fmt.Println("ERROR:", err)
			} else {
				fmt.Println(resultformatter.FormatValue(result))
			}
		}
	}
}
