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

func TestParseNextTokenBinNumber(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("b$0"), TokenNumber, 0.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$1"), TokenNumber, 1.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$1111"), TokenNumber, 15.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$11111111"), TokenNumber, 255.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$111111111111"), TokenNumber, 4095.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$1111111111111111"), TokenNumber, 65535.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$10000000000000000000000000000000"), TokenNumber, -2147483648, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$01111111111111111111111111111111"), TokenNumber, 2147483647, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$11111111111111111111111111111111"), TokenNumber, -1, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("b$111111111111111111111111111111111"), TokenBad, 0, "")
}

func TestParseNextTokenOctNumber(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("o$0"), TokenNumber, 0.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$1"), TokenNumber, 1.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$17"), TokenNumber, 15.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$377"), TokenNumber, 255.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$7777"), TokenNumber, 4095.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$177777"), TokenNumber, 65535.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$20000000000"), TokenNumber, -2147483648, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$17777777777"), TokenNumber, 2147483647, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$37777777777"), TokenNumber, -1, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("o$fffffffff"), TokenBad, 0, "")
}

func TestParseNextTokenHexNumber(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("h$0"), TokenNumber, 0.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$1"), TokenNumber, 1.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$f"), TokenNumber, 15.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$ff"), TokenNumber, 255.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$fff"), TokenNumber, 4095.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$ffff"), TokenNumber, 65535.0, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$80000000"), TokenNumber, -2147483648, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$7fffffff"), TokenNumber, 2147483647, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$ffffffff"), TokenNumber, -1, "")
	assertNextTokenValue(t, CreateLexicalAnalyser("h$fffffffff"), TokenBad, 0, "")
}

func TestParseNextTokenVariable(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("$variablename"), TokenVariable, 0.0, "$variablename")
	assertNextTokenValue(t, CreateLexicalAnalyser("$"), TokenBad, 0.0, "")
}

func TestParseNextTokenIdenitfier(t *testing.T) {
	assertNextTokenValue(t, CreateLexicalAnalyser("identifier"), TokenIdentifier, 0.0, "identifier")
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
