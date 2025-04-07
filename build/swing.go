package build

type Swing struct {
	Kind SwingCMD
}

//go:generate stringer -type SwingCMD --output zz_swing_cmd_string.go
type SwingCMD uint8

const (
	SwingMainHand SwingCMD = 1 + iota
	SwingOffHand
)
