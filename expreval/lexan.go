package expreval

import (
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode"
)

var ErrGeneral = errors.New("general error reading input")
var ErrIdentifierSyntax = errors.New("identifiers must begin with a letter and have a non-zero length")

// Lexical analyser token.
//
//go:generate stringer -type=LexAnToken
type LexAnToken int

const (
	// Error processing another token type.
	TokenBad LexAnToken = iota
	// End of input stream reached.
	TokenEnd
	// Left parentheses "(".
	TokenLParen
	// Right parentheses ")".
	TokenRParen
	// Assign operator "=".
	TokenOpAssign
	// Plus operator "+".
	TokenOpPlus
	// Minus operator "-".
	TokenOpMinus
	// Multiply operator "*".
	TokenOpMultiply
	// Divide operator "/".
	TokenOpDivide
	// Power operator "^".
	TokenOpPower
	// Variable name "$name".
	TokenVariable
	// Number.
	TokenNumber
)

// Lexical Analyser that returns tokens.
type LexicalAnalyser interface {
	// Parse and return the next token.
	ParseNextToken() LexAnToken
	// Gets the current token type.
	GetCurrentToken() LexAnToken
	// Gets the text value of the current token.
	GetTextValue() string
	// Gets the numeric value of the current token.
	GetNumericValue() float64
}

// Lexical Analyser implementation that uses io.Reader.
type LexicalAnalyserReaderImpl struct {
	input        string
	reader       *strings.Reader
	currentToken LexAnToken
	textValue    string
	numericValue float64
}

// Creates a Lexical Analyser for the supplied input.
func CreateLexicalAnalyser(input string) LexicalAnalyser {
	lexAn := new(LexicalAnalyserReaderImpl)
	lexAn.input = input
	lexAn.reader = strings.NewReader(input)
	return lexAn
}

func (lexAn *LexicalAnalyserReaderImpl) ParseNextToken() LexAnToken {

	c, err := nextCharacterIgnoringWhitespace(lexAn.reader)
	if err != nil {
		if err == io.EOF {
			lexAn.currentToken = TokenEnd
		} else {
			lexAn.currentToken = TokenBad
		}

		return lexAn.currentToken
	}

	switch c {
	case '(':
		lexAn.currentToken = TokenLParen
	case ')':
		lexAn.currentToken = TokenRParen
	case '=':
		lexAn.currentToken = TokenOpAssign
	case '+':
		lexAn.currentToken = TokenOpPlus
	case '-':
		lexAn.currentToken = TokenOpMinus
	case '*':
		lexAn.currentToken = TokenOpMultiply
	case '/':
		lexAn.currentToken = TokenOpDivide
	case '^':
		lexAn.currentToken = TokenOpPower
	case '$':
		var identifier, err = parserIdentifier(lexAn.reader)
		if err != nil {
			lexAn.currentToken = TokenBad
			lexAn.textValue = ""
			lexAn.numericValue = 0
		} else {
			lexAn.currentToken = TokenVariable
			lexAn.textValue = "$" + identifier
			lexAn.numericValue = 0
		}
	case '0': // Number.
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		fallthrough
	case '.':
		lexAn.reader.UnreadRune()
		number, err := parseNumber(lexAn.reader)
		if err != nil {
			lexAn.currentToken = TokenBad
			lexAn.textValue = ""
			lexAn.numericValue = 0
		} else {
			lexAn.currentToken = TokenNumber
			lexAn.textValue = ""
			lexAn.numericValue = number
		}
	default:
		lexAn.currentToken = TokenBad
		lexAn.textValue = ""
		lexAn.numericValue = 0
	}

	return lexAn.currentToken
}

func (lexAn *LexicalAnalyserReaderImpl) GetCurrentToken() LexAnToken {
	return lexAn.currentToken
}

func (lexAn *LexicalAnalyserReaderImpl) GetTextValue() string {
	return lexAn.textValue
}

func (lexAn *LexicalAnalyserReaderImpl) GetNumericValue() float64 {
	return lexAn.numericValue
}

func nextCharacterIgnoringWhitespace(reader *strings.Reader) (rune, error) {
	var c rune
	var err error

	for {
		c, _, err = reader.ReadRune()
		if err != nil {
			break
		}

		if !unicode.IsSpace(c) {
			break
		}
	}

	return c, err
}

func parseNumber(reader *strings.Reader) (float64, error) {
	numberString := ""
	foundDecimalPoint := false

	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				reader.UnreadRune()
				break
			} else {
				return 0.0, ErrGeneral
			}
		}

		if c == '.' {
			if foundDecimalPoint {
				reader.UnreadRune()
				break
			} else {
				foundDecimalPoint = true
			}
		} else {
			if !unicode.IsNumber(c) {
				reader.UnreadRune()
				break
			}
		}

		numberString += string(c)
	}

	if len(numberString) == 0 {
		return 0.0, ErrIdentifierSyntax
	}

	return strconv.ParseFloat(numberString, 64)
}

func parserIdentifier(reader *strings.Reader) (string, error) {
	haveFirstLetter := false
	identifer := ""

	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				reader.UnreadRune()
				break
			} else {
				return "", ErrGeneral
			}
		}

		if haveFirstLetter {
			// Identifiers MUST start with a letter.
			if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
				reader.UnreadRune()
				break
			}
		} else {
			if !unicode.IsLetter(c) {
				reader.UnreadRune()
				return "", ErrIdentifierSyntax
			}
		}

		identifer += string(c)
		haveFirstLetter = true
	}

	if len(identifer) == 0 {
		return "", ErrIdentifierSyntax
	}

	return identifer, nil
}
