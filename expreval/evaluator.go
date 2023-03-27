package expreval

import (
	"errors"
	"math"
)

var ErrPrimaryExpected = errors.New("primary expected")
var ErrSyntax = errors.New("syntax error")
var ErrDivideByZero = errors.New("divide by zero")
var ErrMissingClosingParentheses = errors.New("')' expected")
var ErrUnexpectedRightParentheses = errors.New("unexpected ')'")

type Evaluator struct {
	VariableStore map[string]float64
}

func NewEvaluator() *Evaluator {
	evaluator := Evaluator{make(map[string]float64)}
	return &evaluator
}

func (evaluator *Evaluator) Evaluate(expression string) (float64, error) {
	lexAn := CreateLexicalAnalyser(expression)
	result, err := evaluator.getTerm(lexAn, 0, 0)
	if err == nil {
		evaluator.VariableStore["$ans"] = result
	}

	return result, err
}

func (evaluator *Evaluator) getTerm(lexAn LexicalAnalyser, precedence uint, parenthesesLevel uint) (float64, error) {
	// Process the terms at the supplied precedence.
	switch precedence {
	case 0: // OR, RP, END, ERROR.

		// Process any higher operators first.
		leftTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenRParen: // Final exit point (result).
				if parenthesesLevel == 0 { // Check for too many RPs.
					return 0.0, ErrUnexpectedRightParentheses
				}

				return leftTerm, nil

			case TokenEnd: // Final exit point (result).
				return leftTerm, nil

			default: // Systax error in the expression,
				return 0.0, ErrSyntax
			}
		}

	case 1: // PLUS & MINUS.

		// Process any higher operators first.
		leftTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenOpPlus:
				rightTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}
				leftTerm += rightTerm

			case TokenOpMinus:
				rightTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}
				leftTerm -= rightTerm

			default:
				return leftTerm, nil
			}
		}

	case 2: // MULTIPLY, DIVIDE

		// Process any higher operators first.
		leftTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenOpMultiply:
				rightTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}

				leftTerm *= rightTerm

			case TokenOpDivide:
				{
					rightTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
					if err != nil {
						return 0.0, err
					}

					// Prevent a divide by zero.
					if rightTerm == 0.0 {
						return 0.0, ErrDivideByZero
					}

					leftTerm /= rightTerm
				}

			default:
				return leftTerm, nil
			}
		}

	case 3: // POWER.

		// Process any higher operators first.
		leftTerm, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenOpPower:
				p, err := evaluator.getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}

				leftTerm = math.Pow(leftTerm, p)

			default:
				return leftTerm, nil
			}
		}

	default: // Get primary.

		// Process the lexer token.
		switch lexAn.ParseNextToken() {
		case TokenNumber:
			{
				// Store the extracted number.
				v := lexAn.GetNumericValue()

				// Get the next token, so that the token type of the next token is available to the caller of this
				// function.
				lexAn.ParseNextToken()

				// Return the stored number
				return v, nil
			}

		case TokenOpMinus:
			v, err := evaluator.getTerm(lexAn, precedence, parenthesesLevel)
			return -v, err

		case TokenOpPlus:
			return evaluator.getTerm(lexAn, precedence, parenthesesLevel)

		case TokenVariable:
			//// Extract symbol value from global symbol table.
			variableName := lexAn.GetTextValue()
			variableValue := evaluator.VariableStore[variableName]
			var err error = nil

			// Get the next token, so that the token type of the next token is available to the caller of this function.
			// If we have an assign "=" then process the terms after the assign to determine the value of the symbol.
			if lexAn.ParseNextToken() == TokenOpAssign {
				variableValue, err = evaluator.getTerm(lexAn, 0, 0)
				if err == nil {
					evaluator.VariableStore[variableName] = variableValue
				}
			}

			// Return the value of the symbol.
			return variableValue, err

		case TokenLParen:
			{
				// Treat the expression after the parentheses as a new expression and evaluate.
				parenResult, err := evaluator.getTerm(lexAn, 0, parenthesesLevel+1)
				if err != nil {
					return 0.0, err
				}

				// Check expression should have ended on a right parentheses.
				if lexAn.GetCurrentToken() != TokenRParen {
					return 0.0, ErrMissingClosingParentheses
				}

				lexAn.ParseNextToken()

				// Return the value of the expression enclosed by the parentheses.
				return parenResult, nil
			}

		case TokenBad:
			return 0.0, ErrSyntax

		default:
			return 0.0, ErrPrimaryExpected
		}
	}
}
