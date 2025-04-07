package build

type Debuff struct {
	ID      string
	Kind    DebuffCMD
	Casting *Casting
}

//go:generate stringer -type DebuffCMD --output zz_debuff_cmd_string.go
type DebuffCMD uint8

const (
	DebuffRemains DebuffCMD = 1 + iota
	DebuffRemainsExpected
	DebuffCooldownRemains
	DebuffUp
	DebuffDown
	DebuffStack
	DebuffMaxStack
	DebuffAtMaxStack
	DebuffStackPct
	DebuffReact
	DebuffValue
)
