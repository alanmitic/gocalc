package command

import (
	"alanmitic/gocalc/expreval"
)

type Argument struct {
	token        expreval.LexAnToken
	textValue    string
	numericValue float64
}

type Signature []expreval.LexAnToken

type Command interface {
	GetName() string
	GetSignatures() []Signature
	Execute(arguments []Argument) error
	// TODO: Add a GetUsage which returns a string
}
