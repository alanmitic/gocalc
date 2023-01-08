package expreval

import (
	"testing"
)

func TestParseNextTokenOperators(t *testing.T) {
	assertNextToken(t, CreateLexicalAnalyser("("), TokenLParen)
	assertNextToken(t, CreateLexicalAnalyser(")"), TokenRParen)
	assertNextToken(t, CreateLexicalAnalyser("="), TokenOpAssign)
	assertNextToken(t, CreateLexicalAnalyser("+"), TokenOpPlus)
	assertNextToken(t, CreateLexicalAnalyser("-"), TokenOpMinus)
	assertNextToken(t, CreateLexicalAnalyser("*"), TokenOpMultiply)
	assertNextToken(t, CreateLexicalAnalyser("/"), TokenOpDivide)
	assertNextToken(t, CreateLexicalAnalyser("^"), TokenOpPower)
}

func TestParseNextTokenNumber(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("1234567890"), TokenNumber, 1234567890.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("0.123456789"), TokenNumber, 0.123456789, "")
	assertNextTokenValue(t, CreateLexicalAnalyser(".123"), TokenNumber, 0.123, "")
}

func TestParseNextTokenVariable(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("$variablename"), TokenVariable, 0.0, "$variablename")
	assertNextTokenValue(t, CreateLexicalAnalyser("$"), TokenBad, 0.0, "")
}

func TestParseNextTokenMultipleTokens(t *testing.T) {
	lexAn := CreateLexicalAnalyser("()=+-*/^$abcdefg 1234567890")
	assertNextToken(t, lexAn, TokenLParen)
	assertNextToken(t, lexAn, TokenRParen)
	assertNextToken(t, lexAn, TokenOpAssign)
	assertNextToken(t, lexAn, TokenOpPlus)
	assertNextToken(t, lexAn, TokenOpMinus)
	assertNextToken(t, lexAn, TokenOpMultiply)
	assertNextToken(t, lexAn, TokenOpDivide)
	assertNextToken(t, lexAn, TokenOpPower)
	assertNextTokenValue(t, lexAn, TokenVariable, 0.0, "$abcdefg")
	assertNextTokenValue(t, lexAn, TokenNumber, 1234567890.0, "")
	assertNextToken(t, lexAn, TokenEnd)
}

func TestParseNextTokenUnknown(t *testing.T) {
	lexAn := CreateLexicalAnalyser("* £")
	assertNextTokenValue(t, lexAn, TokenOpMultiply, 0.0, "")
	assertNextTokenValue(t, lexAn, TokenBad, 0.0, "")
}

func assertNextToken(t *testing.T, lexAn LexicalAnalyser, expectedToken LexAnToken) {
	token := lexAn.ParseNextToken()
	if token != expectedToken {
		t.Error("Expected:", expectedToken, "Actual:", token)
	}
}

func assertNextTokenValue(t *testing.T, lexAn LexicalAnalyser, expectedToken LexAnToken, expectedNumericValue float64, expectedTextValue string) {
	token := lexAn.ParseNextToken()
	if token != expectedToken {
		t.Error("Expected:", expectedToken, "Actual:", token)
	}

	textValue := lexAn.GetTextValue()
	numericValue := lexAn.GetNumericValue()

	if textValue != expectedTextValue {
		t.Error("Expected:", expectedTextValue, "Actual:", textValue)
	}

	if numericValue != expectedNumericValue {
		t.Error("Expected:", expectedNumericValue, "Actual:", numericValue)
	}
}
