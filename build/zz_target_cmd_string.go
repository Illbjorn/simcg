// Code generated by "stringer -type TargetCMD --output zz_target_cmd_string.go"; DO NOT EDIT.

package build

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TargetLevel-1]
	_ = x[TargetHealth-2]
	_ = x[TargetAdds-3]
	_ = x[TargetAddsNever-4]
	_ = x[TargetDistance-5]
	_ = x[TargetCurrentTarget-6]
	_ = x[TargetName-7]
	_ = x[TargetTTD-8]
}

const _TargetCMD_name = "TargetLevelTargetHealthTargetAddsTargetAddsNeverTargetDistanceTargetCurrentTargetTargetNameTargetTTD"

var _TargetCMD_index = [...]uint8{0, 11, 23, 33, 48, 62, 81, 91, 100}

func (i TargetCMD) String() string {
	i -= 1
	if i >= TargetCMD(len(_TargetCMD_index)-1) {
		return "TargetCMD(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TargetCMD_name[_TargetCMD_index[i]:_TargetCMD_index[i+1]]
}
