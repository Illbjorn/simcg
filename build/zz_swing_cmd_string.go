// Code generated by "stringer -type SwingCMD --output zz_swing_cmd_string.go"; DO NOT EDIT.

package build

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SwingMainHand-1]
	_ = x[SwingOffHand-2]
}

const _SwingCMD_name = "SwingMainHandSwingOffHand"

var _SwingCMD_index = [...]uint8{0, 13, 25}

func (i SwingCMD) String() string {
	i -= 1
	if i >= SwingCMD(len(_SwingCMD_index)-1) {
		return "SwingCMD(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SwingCMD_name[_SwingCMD_index[i]:_SwingCMD_index[i+1]]
}
