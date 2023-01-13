package expreval

import (
	"errors"
	"math"
)

var PrimaryExpectedError = errors.New("Primary expected")
var SyntaxError = errors.New("Syntax error")
var DivideByZeroError = errors.New("Divide by zero")
var MissingClosingParentheses = errors.New("')' expected")
var UnexpectedRightParentheses = errors.New("unexpected ')'")

type Evaluator struct {
}

func (evaluator *Evaluator) Evaluate(expression string) (float64, error) {
	lexAn := CreateLexicalAnalyser(expression)
	result, err := getTerm(lexAn, 0, 0)
	return result, err
}

func getTerm(lexAn LexicalAnalyser, precedence uint, parenthesesLevel uint) (float64, error) {
	// Process the terms at the supplied precedence.
	switch precedence {
	case 0: // OR, RP, END, ERROR.

		// Process any higher operators first.
		leftTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenRParen: // Final exit point (result).
				if parenthesesLevel == 0 { // Check for too many RPs.
					return 0.0, UnexpectedRightParentheses
				}

				return leftTerm, nil

			case TokenEnd: // Final exit point (result).
				return leftTerm, nil

			default: // Systax error in the expression,
				return 0.0, SyntaxError
			}
		}

	case 1: // PLUS & MINUS.

		// Process any higher operators first.
		leftTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenOpPlus:
				rightTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}
				leftTerm += rightTerm
				break

			case TokenOpMinus:
				rightTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}
				leftTerm -= rightTerm
				break

			default:
				return leftTerm, nil
			}
		}

	case 2: // MULTIPLY, DIVIDE

		// Process any higher operators first.
		leftTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenOpMultiply:
				rightTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}

				leftTerm *= rightTerm
				break

			case TokenOpDivide:
				{
					rightTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
					if err != nil {
						return 0.0, err
					}

					// Prevent a divide by zero.
					if rightTerm == 0.0 {
						return 0.0, DivideByZeroError
					}

					leftTerm /= rightTerm
				}
				break

			default:
				return leftTerm, nil
			}
		}

	case 3: // POWER.

		// Process any higher operators first.
		leftTerm, err := getTerm(lexAn, precedence+1, parenthesesLevel)
		if err != nil {
			return 0.0, err
		}

		for {
			// Process the lexer token.
			switch lexAn.GetCurrentToken() {
			case TokenOpPower:
				p, err := getTerm(lexAn, precedence+1, parenthesesLevel)
				if err != nil {
					return 0.0, err
				}

				leftTerm = math.Pow(leftTerm, p)
				break

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
			v, err := getTerm(lexAn, precedence, parenthesesLevel)
			return -v, err

		case TokenOpPlus:
			return getTerm(lexAn, precedence, parenthesesLevel)

		case TokenLParen:
			{
				// Treat the expression after the parentheses as a new expression and evaluate.
				parenResult, err := getTerm(lexAn, 0, parenthesesLevel+1)
				if err != nil {
					return 0.0, err
				}

				// Check expression should have ended on a right parentheses.
				if lexAn.GetCurrentToken() != TokenRParen {
					return 0.0, MissingClosingParentheses
				}

				lexAn.ParseNextToken()

				// Return the value of the expression enclosed by the parentheses.
				return parenResult, nil
			}

		case TokenBad:
			return 0.0, SyntaxError

		default:
			return 0.0, PrimaryExpectedError
		}
	}
}
