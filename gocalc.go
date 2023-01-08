package main

import (
	"alanmitic/gocalc/expreval"
	"fmt"
)

func main() {
	fmt.Println("Go Calc")
	result := expreval.Evaluate("abc")
	fmt.Println("result = ", result)
}
