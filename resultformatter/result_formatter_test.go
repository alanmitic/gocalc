package resultformatter

import (
	"testing"
)

func TestResultFormatterDefault(t *testing.T) {
	resultFormatter := NewResultFormatter()

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeReal {
		t.Error("Expected:", OutputModeReal, "Actual:", actualOutputMode)
	}

	actualPrecision := resultFormatter.GetPrecision()
	if actualPrecision != -1 {
		t.Error("Expected:", -1, "Actual:", actualPrecision)
	}
}

func TestResultFormatterSetOutputModeFixed(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeFixed)

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeFixed {
		t.Error("Expected:", OutputModeFixed, "Actual:", actualOutputMode)
	}
}

func TestResultFormatterSetOutputModeReal(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeReal)

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeReal {
		t.Error("Expected:", OutputModeReal, "Actual:", actualOutputMode)
	}
}

func TestResultFormatterSetOutputModeScientific(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeScientific)

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeScientific {
		t.Error("Expected:", OutputModeScientific, "Actual:", actualOutputMode)
	}
}

func TestResultFormatterSetOutputModeBinary(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeBinary)

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeBinary {
		t.Error("Expected:", OutputModeBinary, "Actual:", actualOutputMode)
	}
}

func TestResultFormatterSetOutputModeOctal(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeOctal)

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeOctal {
		t.Error("Expected:", OutputModeOctal, "Actual:", actualOutputMode)
	}
}

func TestResultFormatterSetOutputModeHexadecimal(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeHexadecimal)

	actualOutputMode := resultFormatter.GetOutputMode()
	if actualOutputMode != OutputModeHexadecimal {
		t.Error("Expected:", OutputModeHexadecimal, "Actual:", actualOutputMode)
	}
}

func TestResultFormatterSetPrecision(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetPrecision(6)

	actualPrecision := resultFormatter.GetPrecision()
	if actualPrecision != 6 {
		t.Error("Expected:", 6, "Actual:", actualPrecision)
	}
}

func TestResultFormatterFormatValueDefault(t *testing.T) {
	resultFormatter := NewResultFormatter()
	assertFormattedValue(t, resultFormatter, 1234.56789, "1234.56789")
}

func TestResultFormatterFormatValueFixed(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeFixed)
	resultFormatter.SetPrecision(3)
	assertFormattedValue(t, resultFormatter, 1234.56789, "1234.568")
}

func TestResultFormatterFormatValueScientific(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeScientific)
	resultFormatter.SetPrecision(3)
	assertFormattedValue(t, resultFormatter, 1234.56789, "1.235e+03")
}

func TestResultFormatterFormatValueBinary(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeBinary)
	assertFormattedValue(t, resultFormatter, 123456789, "00000111010110111100110100010101")
}

func TestResultFormatterFormatValueOctal(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeOctal)
	assertFormattedValue(t, resultFormatter, 123456789, "000726746425")
}

func TestResultFormatterFormatHexadecimal(t *testing.T) {
	resultFormatter := NewResultFormatter()
	resultFormatter.SetOutputMode(OutputModeHexadecimal)
	assertFormattedValue(t, resultFormatter, 123456789, "075bcd15")
}

func assertFormattedValue(t *testing.T, resultFormatter ResultFormatter, inputValue float64, expectedFormattedValue string) {
	formattedValue := resultFormatter.FormatValue(inputValue)
	if formattedValue != expectedFormattedValue {
		t.Error("Expected:", expectedFormattedValue, "Actual:", formattedValue)
	}
}
