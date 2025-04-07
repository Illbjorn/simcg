package build

type RaidEvent struct {
	Kind   RaidEventCMD
	Filter RaidEventFilterCMD
}

//go:generate stringer -type RaidEventCMD --output zz_raid_event_cmd_string.go
type RaidEventCMD uint8

const (
	RaidEventAdds RaidEventCMD = 1 + iota
	RaidEventCasting
	RaidEventDamageTaken
	RaidEventDamage
	RaidEventDistraction
	RaidEventFlying
	RaidEventHeal
	RaidEventInterrupt
	RaidEventInvulnerable
	RaidEventMovement
	RaidEventPositionSwitch
	RaidEventStun
	RaidEventVulnerable
)

//go:generate stringer -type RaidEventFilterCMD --output zz_raid_event_filter_cmd_string.go
type RaidEventFilterCMD uint8

const (
	RaidEventFilterIn RaidEventFilterCMD = 1 + iota
	RaidEventFilterAmount
	RaidEventFilterDuration
	RaidEventFilterCooldown
	RaidEventFilterExists
	RaidEventFilterDistance
	RaidEventFilterMaxDistance
	RaidEventFilterMinDistance
	RaidEventFilterToPct
	RaidEventFilterUp
	RaidEventFilterRemains
)
