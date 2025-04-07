package build

type Buff struct {
	ID   string
	Kind BuffCMD
}

//go:generate stringer -type BuffCMD --output zz_buff_cmd_string.go
type BuffCMD uint8

const (
	BuffRemains BuffCMD = 1 + iota
	BuffRemainsExpected
	BuffCooldownRemains
	BuffUp
	BuffDown
	BuffStack
	BuffMaxStack
	BuffAtMaxStack
	BuffStackPct
	BuffReact
	BuffValue
)
