package build

type Cooldown struct {
	ID   string
	Kind CooldownCMD
}

//go:generate stringer -type CooldownCMD --output zz_cooldown_cmd_string.go
type CooldownCMD uint8

const (
	CooldownDuration CooldownCMD = 1 + iota
	CooldownRemains
	CooldownUp
	CooldownReady
	CooldownCharges
	CooldownChargesFractional
	CooldownFullRechargeTime
	CooldownMaxCharges
	CooldownDurationExpected
	CooldownRemainsExpected
)
