// Code generated by "stringer -type=ARPAllTargets -linecomment"; DO NOT EDIT.

package nethelpers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ARPAllTargetsAny-0]
	_ = x[ARPAllTargetsAll-1]
}

const _ARPAllTargets_name = "anyall"

var _ARPAllTargets_index = [...]uint8{0, 3, 6}

func (i ARPAllTargets) String() string {
	if i >= ARPAllTargets(len(_ARPAllTargets_index)-1) {
		return "ARPAllTargets(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ARPAllTargets_name[_ARPAllTargets_index[i]:_ARPAllTargets_index[i+1]]
}
