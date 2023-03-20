package command

import (
	"alanmitic/gocalc/resultformatter"
	"testing"
)

func assertNilCommandAndError(t *testing.T, command Command, actualError error, expectedError error) {
	if command != nil {
		t.Error("Expected:", nil, "Actual:", command.GetName())
	}
	if actualError != expectedError {
		t.Error("Expected:", expectedError, "Actual:", actualError)
	}
}

func assertCommand(t *testing.T, command Command, expectedCommandName string) {
	commandName := command.GetName()
	if commandName != expectedCommandName {
		t.Error("Expected:", expectedCommandName, "Actual:", commandName)
	}
}

func assertArguments(t *testing.T, arguments []Argument, expectedArguments []Argument) {
	for index, expectedArgument := range expectedArguments {
		if arguments[index].token != expectedArgument.token {
			t.Error("Expected:", expectedArgument.token, "Actual:", arguments[index].token)
		}
		if arguments[index].textValue != expectedArgument.textValue {
			t.Error("Expected:", expectedArgument.textValue, "Actual:", arguments[index].textValue)
		}
		if arguments[index].numericValue != expectedArgument.numericValue {
			t.Error("Expected:", expectedArgument.numericValue, "Actual:", arguments[index].numericValue)
		}
	}
}

func assertOutputModeAndPrecision(t *testing.T, resultFormatter resultformatter.ResultFormatter,
	expectedOutputMode resultformatter.OutputMode, expectedPrecision int) {
	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != expectedOutputMode {
		t.Error("Expected:", expectedOutputMode, "Actual:", actualOutputMode)
	}

	actualPrecision := resultFormatter.GetPrecision()
	if actualPrecision != expectedPrecision {
		t.Error("Expected:", expectedPrecision, "Actual:", actualPrecision)
	}
}
