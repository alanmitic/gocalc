package expreval

import (
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var ErrGeneral = errors.New("general error reading input")
var ErrIdentifierSyntax = errors.New("identifiers must begin with a letter and have a non-zero length")
var ErrOverflow = errors.New("value too large")

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
	// Identifier.
	TokenIdentifier
)

//go:generate stringer -type=BaseModifier
type BaseModifier int

const (
	BaseModifierNone BaseModifier = iota
	BaseModifierBin
	BaseModifierOct
	BaseModifierHex
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

var binRangeTable = &unicode.RangeTable{
	R16:         []unicode.Range16{{0x0030, 0x0031, 1}},
	R32:         []unicode.Range32{},
	LatinOffset: 1,
}

var octRangeTable = &unicode.RangeTable{
	R16:         []unicode.Range16{{0x0030, 0x0037, 1}},
	R32:         []unicode.Range32{},
	LatinOffset: 1,
}

var hexRangeTable = &unicode.RangeTable{
	R16:         []unicode.Range16{{0x0030, 0x0039, 1}, {0x0051, 0x005A, 1}, {0x0061, 0x007A, 1}},
	R32:         []unicode.Range32{},
	LatinOffset: 3,
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
		var identifier, err = parserIdentifier(lexAn.reader, "")
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
		number, err := parseNumber(lexAn.reader, BaseModifierNone)
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
		// An idenitifier or bad token.
		if unicode.IsLetter(c) {
			symbolChar, _, _ := lexAn.reader.ReadRune()

			preReadfragment := string(c) + string(symbolChar)
			baseModifier := extractNumberBaseModifier(preReadfragment)
			if baseModifier != BaseModifierNone {
				// Handle b$n, o%n or h$n
				number, err := parseNumber(lexAn.reader, baseModifier)
				if err != nil {
					lexAn.currentToken = TokenBad
					lexAn.textValue = ""
					lexAn.numericValue = 0
				} else {
					lexAn.currentToken = TokenNumber
					lexAn.textValue = ""
					lexAn.numericValue = number
				}
			} else {
				identifier, err := parserIdentifier(lexAn.reader, preReadfragment)
				if err != nil {
					lexAn.currentToken = TokenBad
					lexAn.textValue = ""
					lexAn.numericValue = 0
				} else {
					lexAn.currentToken = TokenIdentifier
					lexAn.textValue = identifier
					lexAn.numericValue = 0
				}
			}
		} else {
			lexAn.currentToken = TokenBad
			lexAn.textValue = ""
			lexAn.numericValue = 0
		}
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

func parseNumber(reader *strings.Reader, baseModifier BaseModifier) (float64, error) {
	numberString := ""
	foundDecimalPoint := false

	rangeTable := unicode.Digit
	baseSupportsDecimalPoint := true

	switch baseModifier {
	case BaseModifierBin:
		rangeTable = binRangeTable
		baseSupportsDecimalPoint = false
	case BaseModifierOct:
		rangeTable = octRangeTable
		baseSupportsDecimalPoint = false
	case BaseModifierHex:
		rangeTable = hexRangeTable
		baseSupportsDecimalPoint = false
	}

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

		if c == '.' && baseSupportsDecimalPoint {
			if foundDecimalPoint {
				reader.UnreadRune()
				break
			} else {
				foundDecimalPoint = true
			}
		} else {
			if !unicode.In(c, rangeTable) {
				reader.UnreadRune()
				break
			}
		}

		numberString += string(c)
	}

	if len(numberString) == 0 {
		return 0.0, ErrIdentifierSyntax
	}

	if baseModifier != BaseModifierNone {
		var base int
		switch baseModifier {
		case BaseModifierBin:
			base = 2
		case BaseModifierOct:
			base = 8
		case BaseModifierHex:
			base = 16
		}

		intNumber64, err := strconv.ParseInt(numberString, base, 64)
		if err != nil {
			return 0.0, err
		}

		if intNumber64 > math.MaxUint32 {
			return 0.0, ErrOverflow
		}

		truncatedInt32 := int32(intNumber64)
		return float64(truncatedInt32), err
	} else {
		return strconv.ParseFloat(numberString, 64)
	}
}

func extractNumberBaseModifier(modifier string) BaseModifier {
	switch modifier {
	case "b$":
		return BaseModifierBin
	case "o$":
		return BaseModifierOct
	case "h$":
		return BaseModifierHex
	default:
		return BaseModifierNone
	}
}

func parserIdentifier(reader *strings.Reader, preReadFragment string) (string, error) {
	haveFirstLetter := false
	identifer := preReadFragment

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
