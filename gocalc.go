package main

import (
	"alanmitic/gocalc/expreval"
	"fmt"
)

func main() {
	fmt.Println("Go Calc")
	evaluator := new(expreval.Evaluator)

	result, err := evaluator.Evaluate("1 + 3 * 4")
	fmt.Println("result = ", result)
	fmt.Println("err = ", err)
}
