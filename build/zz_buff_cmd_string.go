// Code generated by "stringer -type BuffCMD --output zz_buff_cmd_string.go"; DO NOT EDIT.

package build

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BuffRemains-1]
	_ = x[BuffRemainsExpected-2]
	_ = x[BuffCooldownRemains-3]
	_ = x[BuffUp-4]
	_ = x[BuffDown-5]
	_ = x[BuffStack-6]
	_ = x[BuffMaxStack-7]
	_ = x[BuffAtMaxStack-8]
	_ = x[BuffStackPct-9]
	_ = x[BuffReact-10]
	_ = x[BuffValue-11]
}

const _BuffCMD_name = "BuffRemainsBuffRemainsExpectedBuffCooldownRemainsBuffUpBuffDownBuffStackBuffMaxStackBuffAtMaxStackBuffStackPctBuffReactBuffValue"

var _BuffCMD_index = [...]uint8{0, 11, 30, 49, 55, 63, 72, 84, 98, 110, 119, 128}

func (i BuffCMD) String() string {
	i -= 1
	if i >= BuffCMD(len(_BuffCMD_index)-1) {
		return "BuffCMD(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _BuffCMD_name[_BuffCMD_index[i]:_BuffCMD_index[i+1]]
}
