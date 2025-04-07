package build

type Trinket struct {
	ID       string
	Slot     TrinketSlot
	Kind     TrinketCMD
	Resource Resource
}

//go:generate stringer -type TrinketSlot --output zz_trinket_slot_string.go
type TrinketSlot uint8

const (
	TrinketSlot1 TrinketSlot = 1 + iota
	TrinketSlot2
)

//go:generate stringer -type TrinketCMD --output zz_trinket_cmd_string.go
type TrinketCMD uint8

const (
	TrinketCooldown TrinketCMD = 1 + iota
	TrinketIs
	TrinketReadyCooldown
	TrinketHasUseBuff
	TrinketBuff
	TrinketHasStat
	TrinketCastTime
	TrinketHasCooldown
	TrinketRemains
	TrinketProc
	TrinketDuration
)
