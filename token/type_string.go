// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Whitespace-1]
	_ = x[True-2]
	_ = x[False-3]
	_ = x[Null-4]
	_ = x[Number-5]
	_ = x[String-6]
}

const _Type_name = "WhitespaceTrueFalseNullNumberString"

var _Type_index = [...]uint8{0, 10, 14, 19, 23, 29, 35}

func (i Type) String() string {
	i -= 1
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}