package command

import (
	"alanmitic/gocalc/expreval"
	"alanmitic/gocalc/resultformatter"
	"errors"
)

var ErrGeneral = errors.New("general error parsing command")
var ErrTooManyArgs = errors.New("too many arguments")
var ErrInvalidArgs = errors.New("command arguments are invalid")
var ErrNotFound = errors.New("command not found")

type CommandParser struct {
	commands map[string]Command
}

func NewCommandParser(evaluator *expreval.Evaluator, resultformatter resultformatter.ResultFormatter) *CommandParser {
	commandParser := CommandParser{}

	commandParser.commands = make(map[string]Command)

	addCommand(commandParser.commands, NewCommandExit())
	addCommand(commandParser.commands, NewCommandFix(resultformatter))
	addCommand(commandParser.commands, NewCommandReal(resultformatter))
	addCommand(commandParser.commands, NewCommandSci(resultformatter))
	addCommand(commandParser.commands, NewCommandBin(resultformatter))
	addCommand(commandParser.commands, NewCommandOct(resultformatter))
	addCommand(commandParser.commands, NewCommandHex(resultformatter))
	addCommand(commandParser.commands, NewCommandVars(evaluator))
	addCommand(commandParser.commands, NewCommandHelp(commandParser.commands))

	return &commandParser
}

func addCommand(commands map[string]Command, command Command) {
	commands[command.GetName()] = command
}

func (commandParser *CommandParser) ParseCommand(input string) (Command, []Argument, error) {
	var command Command = nil
	lexAn := expreval.CreateLexicalAnalyser(input)

	// Get first token and it needs to be an identifier for it to be a command.
	token := lexAn.ParseNextToken()
	if token == expreval.TokenIdentifier {
		// Find the command.
		command = commandParser.commands[lexAn.GetTextValue()]
		if command != nil {
			arguments, err := parseArguments(lexAn, command)
			if err != nil {
				return nil, nil, err
			}
			return command, arguments, nil
		} else {
			return nil, nil, ErrNotFound
		}
	}

	return nil, nil, nil
}

func parseArguments(lexAn expreval.LexicalAnalyser, command Command) ([]Argument, error) {
	signatures := command.GetSignatures()
	maxNumArgs := 0
	for _, signature := range signatures {
		numArgs := len(signature)
		if numArgs > maxNumArgs {
			maxNumArgs = numArgs
		}
	}

	// Slurp the rest of the tokens as arguments.
	arguments, err := slurpArguments(lexAn, maxNumArgs)
	if err != nil {
		return nil, err
	}

	// Find the matching signature.
	findMatchingSignature(signatures, arguments)

	// If no signature is found return an error.

	return arguments, nil
}

func slurpArguments(lexAn expreval.LexicalAnalyser, maxNumArgs int) ([]Argument, error) {
	//arguments := make([]Argument, maxNumArgs)
	arguments := []Argument{}
	numArgs := 0

	for {
		token := lexAn.ParseNextToken()

		if token == expreval.TokenBad {
			return nil, ErrGeneral
		}

		if token == expreval.TokenEnd {
			break
		}

		numArgs++
		if numArgs > maxNumArgs {
			return nil, ErrTooManyArgs
		}

		arguments = append(arguments, Argument{token, lexAn.GetTextValue(), lexAn.GetNumericValue()})
	}

	return arguments, nil
}

func findMatchingSignature(signatures []Signature, arguments []Argument) (Signature, error) {

	for _, signature := range signatures {
		if len(signature) != len(arguments) {
			break
		}

		matchedSignature := signature
		for tokenIndex, signatureToken := range signature {
			if signatureToken != arguments[tokenIndex].token {
				matchedSignature = nil
				break
			}
		}

		if matchedSignature != nil {
			return matchedSignature, nil
		}
	}

	return nil, ErrInvalidArgs
}
