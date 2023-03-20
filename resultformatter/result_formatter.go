package resultformatter

import (
	"fmt"
	"strconv"
)

//go:generate stringer -type=OutputMode
type OutputMode int

const (
	OutputModeFixed OutputMode = iota
	OutputModeReal
	OutputModeScientific
	OutputModeBinary
	OutputModeOctal
	OutputModeHexadecimal
)

type ResultFormatter interface {
	GetOutputMode() OutputMode
	SetOutputMode(outputMode OutputMode)
	GetPrecision() int
	SetPrecision(precision int)
	FormatValue(value float64) string
}

type ResultFormatterImpl struct {
	outputMode OutputMode
	precision  int
}

func NewResultFormatter() ResultFormatter {
	resultFormatter := new(ResultFormatterImpl)
	resultFormatter.outputMode = OutputModeReal
	resultFormatter.precision = -1
	return resultFormatter
}

func (resultFormatter *ResultFormatterImpl) GetOutputMode() OutputMode {
	return resultFormatter.outputMode
}

func (resultFormatter *ResultFormatterImpl) SetOutputMode(outputMode OutputMode) {
	resultFormatter.outputMode = outputMode
}

func (resultFormatter *ResultFormatterImpl) GetPrecision() int {
	return resultFormatter.precision
}

func (resultFormatter *ResultFormatterImpl) SetPrecision(precision int) {
	resultFormatter.precision = precision
}

func (resultFormatter *ResultFormatterImpl) FormatValue(value float64) string {
	formattedValue := ""

	switch resultFormatter.outputMode {
	case OutputModeFixed:
		formattedValue = strconv.FormatFloat(value, 'f', resultFormatter.precision, 64)
	case OutputModeReal:
		formattedValue = strconv.FormatFloat(value, 'g', resultFormatter.precision, 64)
	case OutputModeScientific:
		formattedValue = strconv.FormatFloat(value, 'e', resultFormatter.precision, 64)
	case OutputModeBinary:
		intValue := int(value)
		formattedValue = fmt.Sprintf("%032b", intValue)
	case OutputModeOctal:
		intValue := int(value)
		formattedValue = fmt.Sprintf("%012o", intValue)
	case OutputModeHexadecimal:
		intValue := int(value)
		formattedValue = fmt.Sprintf("%08x", intValue)
	}

	return formattedValue
}
