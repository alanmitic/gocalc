// Code generated by "stringer -type=OutputMode"; DO NOT EDIT.

package resultformatter

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OutputModeFixed-0]
	_ = x[OutputModeReal-1]
	_ = x[OutputModeScientific-2]
	_ = x[OutputModeBinary-3]
	_ = x[OutputModeOctal-4]
	_ = x[OutputModeHexadecimal-5]
}

const _OutputMode_name = "OutputModeFixedOutputModeRealOutputModeScientificOutputModeBinaryOutputModeOctalOutputModeHexadecimal"

var _OutputMode_index = [...]uint8{0, 15, 29, 49, 65, 80, 101}

func (i OutputMode) String() string {
	if i < 0 || i >= OutputMode(len(_OutputMode_index)-1) {
		return "OutputMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OutputMode_name[_OutputMode_index[i]:_OutputMode_index[i+1]]
}
