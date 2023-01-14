package main

import (
	"alanmitic/gocalc/expreval"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	evaluator := expreval.NewEvaluator()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("gocalc version 0.0.1\n\nType 'exit' and ENTER to quit.")

	for {
		fmt.Print("gocalc >> ")
		exprOrCmd, _ := reader.ReadString('\n')
		if strings.TrimSpace(exprOrCmd) == "exit" {
			return
		}

		result, err := evaluator.Evaluate(exprOrCmd)
		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Println(result)
		}
	}
}
