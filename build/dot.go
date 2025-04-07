package build

type DOT struct {
	ID   string
	Kind DOTCMD
}

//go:generate stringer -type DOTCMD --output zz_dot_cmd_string.go
type DOTCMD uint8

const (
	DotDuration DOTCMD = 1 + iota
	DotUp
	DotDown
	DotRemainsExpected
	DotModifier
	DotRemains
	DotRefreshable
	DotTicking
	DotTicksAdded
	DotTickDmg
	DotTicksRemain
	DotSpellPower
	DotAttackPower
	DotMultiplier
	DotHastePct
	DotCurrentTicks
	DotTicks
	DotCritPct
	DotCritDmg
	DotTickTimeRemains
)
