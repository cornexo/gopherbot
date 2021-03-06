// Code generated by "stringer -type=pipelineType constants.go"; DO NOT EDIT.

package bot

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[unset-0]
	_ = x[plugCommand-1]
	_ = x[plugMessage-2]
	_ = x[catchAll-3]
	_ = x[jobTrigger-4]
	_ = x[spawnedTask-5]
	_ = x[scheduled-6]
	_ = x[jobCmd-7]
}

const _pipelineType_name = "unsetplugCommandplugMessagecatchAlljobTriggerspawnedTaskscheduledjobCmd"

var _pipelineType_index = [...]uint8{0, 5, 16, 27, 35, 45, 56, 65, 71}

func (i pipelineType) String() string {
	if i < 0 || i >= pipelineType(len(_pipelineType_index)-1) {
		return "pipelineType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _pipelineType_name[_pipelineType_index[i]:_pipelineType_index[i+1]]
}
