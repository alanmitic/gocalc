package expreval

import (
	"testing"
)

func TestEvaluateAdd(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("1+2+3+4 + 5 + 6+7+8+9")
	assertEvaluatedResult(t, 45, nil, result, err)
}

func TestEvaluateMinus(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("9-8-7  -  6 -5-4-3-2-1-0")
	assertEvaluatedResult(t, -27, nil, result, err)
}

func TestEvaluateMultiply(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("9*8*7  *  6 *5*4*3*2*1")
	assertEvaluatedResult(t, 362880, nil, result, err)
}

func TestEvaluateDivide(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("10000/500/4  /  2 /1")
	assertEvaluatedResult(t, 2.5, nil, result, err)
}

func TestEvaluateDivideByZero(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("100 / 0")
	assertEvaluatedResult(t, 0.0, ErrDivideByZero, result, err)
}

func TestEvaluatePower(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("2 ^ 4")
	assertEvaluatedResult(t, 16, nil, result, err)
}

func TestEvaluatePrimaryPlus(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("  +123   ")
	assertEvaluatedResult(t, 123, nil, result, err)
}

func TestEvaluatePrimaryMinus(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("   -123   ")
	assertEvaluatedResult(t, -123, nil, result, err)
}

func TestEvaluatePrimaryNumber(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("   1234567890   ")
	assertEvaluatedResult(t, 1234567890, nil, result, err)
}

func TestEvaluateParenthesis(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("2 * (3 + (1 / 2))")
	assertEvaluatedResult(t, 7, nil, result, err)
}

func TestEvaluateUnclosedParenthesis(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("2 * (3 + (1 / 2)")
	assertEvaluatedResult(t, 0.0, ErrMissingClosingParentheses, result, err)
}

func TestEvaluateTooManyParenthesis(t *testing.T) {
	evaluator := NewEvaluator()
	result, err := evaluator.Evaluate("2 * (3 + (1 / 2)))")
	assertEvaluatedResult(t, 0.0, ErrUnexpectedRightParentheses, result, err)
}

func TestEvaluateVariableAns(t *testing.T) {
	evaluator := NewEvaluator()
	evaluator.Evaluate("100 + 200")
	assertVariableValue(t, evaluator, "$ans", 300)
}

func TestEvaluateAssignVariable(t *testing.T) {
	evaluator := NewEvaluator()
	evaluator.Evaluate("$a = 1 + 2 + 3")
	assertVariableValue(t, evaluator, "$ans", 6)
	assertVariableValue(t, evaluator, "$a", 6)
}

func TestEvaluateUseVariable(t *testing.T) {
	evaluator := NewEvaluator()
	evaluator.Evaluate("$a = 1 + 2 + 3")
	evaluator.Evaluate("$b = $a + 3")
	assertVariableValue(t, evaluator, "$a", 6)
	assertVariableValue(t, evaluator, "$b", 9)
	assertVariableValue(t, evaluator, "$ans", 9)
}

func assertVariableValue(t *testing.T, evaluator *Evaluator, variableName string, expectedvariableValue float64) {
	actualVariableValue, variableFound := evaluator.VariableStore[variableName]

	if !variableFound {
		t.Error("Variable not found:", variableName)
	}

	if actualVariableValue != expectedvariableValue {
		t.Error("Variable: ", variableName, " Expected:", expectedvariableValue, "Actual:", actualVariableValue)
	}
}

func assertEvaluatedResult(t *testing.T, expectedResult float64, expectedError error, actualResult float64, actualError error) {
	if actualResult != expectedResult {
		t.Error("Expected:", expectedResult, "Actual:", actualResult)
	}

	if actualError != expectedError {
		t.Error("Expected:", expectedError, "Actual:", actualError)
	}
}
