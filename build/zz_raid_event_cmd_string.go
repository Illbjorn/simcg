// Code generated by "stringer -type RaidEventCMD --output zz_raid_event_cmd_string.go"; DO NOT EDIT.

package build

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RaidEventAdds-1]
	_ = x[RaidEventCasting-2]
	_ = x[RaidEventDamageTaken-3]
	_ = x[RaidEventDamage-4]
	_ = x[RaidEventDistraction-5]
	_ = x[RaidEventFlying-6]
	_ = x[RaidEventHeal-7]
	_ = x[RaidEventInterrupt-8]
	_ = x[RaidEventInvulnerable-9]
	_ = x[RaidEventMovement-10]
	_ = x[RaidEventPositionSwitch-11]
	_ = x[RaidEventStun-12]
	_ = x[RaidEventVulnerable-13]
}

const _RaidEventCMD_name = "RaidEventAddsRaidEventCastingRaidEventDamageTakenRaidEventDamageRaidEventDistractionRaidEventFlyingRaidEventHealRaidEventInterruptRaidEventInvulnerableRaidEventMovementRaidEventPositionSwitchRaidEventStunRaidEventVulnerable"

var _RaidEventCMD_index = [...]uint8{0, 13, 29, 49, 64, 84, 99, 112, 130, 151, 168, 191, 204, 223}

func (i RaidEventCMD) String() string {
	i -= 1
	if i >= RaidEventCMD(len(_RaidEventCMD_index)-1) {
		return "RaidEventCMD(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RaidEventCMD_name[_RaidEventCMD_index[i]:_RaidEventCMD_index[i+1]]
}
