// Code generated by "stringer -type=Event"; DO NOT EDIT.

package bot

import "strconv"

const _Event_name = "IgnoredUserBotDirectMessageAdminCheckPassedAdminCheckFailedMultipleMatchesNoActionAuthNoRunMisconfiguredAuthNoRunPlugNotAvailableAuthRanSuccessAuthRanFailAuthRanMechanismFailedAuthRanFailNormalAuthRanFailOtherAuthNoRunNotFoundElevNoRunMisconfiguredElevNoRunNotAvailableElevRanSuccessElevRanFailElevRanMechanismFailedElevRanFailNormalElevRanFailOtherElevNoRunNotFoundCommandTaskRanAmbientTaskRanCatchAllsRanCatchAllTaskRanTriggeredTaskRanScheduledTaskRanRunJobTaskRanGoPluginRanScriptPluginBadPathScriptPluginBadInterpreterScriptTaskRanScriptPluginStderrOutputScriptPluginErrExit"

var _Event_index = [...]uint16{0, 11, 27, 43, 59, 82, 104, 129, 143, 154, 176, 193, 209, 226, 248, 269, 283, 294, 316, 333, 349, 366, 380, 394, 406, 421, 437, 453, 466, 477, 496, 522, 535, 559, 578}

func (i Event) String() string {
	if i < 0 || i >= Event(len(_Event_index)-1) {
		return "Event(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Event_name[_Event_index[i]:_Event_index[i+1]]
}
