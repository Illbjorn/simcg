// Code generated by "stringer -type BaseValueKind --output zz_base_value_kind_string.go"; DO NOT EDIT.

package build

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BaseID-1]
	_ = x[BaseNum-2]
}

const _BaseValueKind_name = "BaseIDBaseNum"

var _BaseValueKind_index = [...]uint8{0, 6, 13}

func (i BaseValueKind) String() string {
	i -= 1
	if i >= BaseValueKind(len(_BaseValueKind_index)-1) {
		return "BaseValueKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _BaseValueKind_name[_BaseValueKind_index[i]:_BaseValueKind_index[i+1]]
}
