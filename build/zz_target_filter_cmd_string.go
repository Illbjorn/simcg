// Code generated by "stringer -type TargetFilterCMD --output zz_target_filter_cmd_string.go"; DO NOT EDIT.

package build

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TargetFilterPct-1]
}

const _TargetFilterCMD_name = "TargetFilterPct"

var _TargetFilterCMD_index = [...]uint8{0, 15}

func (i TargetFilterCMD) String() string {
	i -= 1
	if i >= TargetFilterCMD(len(_TargetFilterCMD_index)-1) {
		return "TargetFilterCMD(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TargetFilterCMD_name[_TargetFilterCMD_index[i]:_TargetFilterCMD_index[i+1]]
}
