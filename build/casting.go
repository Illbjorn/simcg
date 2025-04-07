package build

type Casting struct {
	Kind CastingCMD
}

//go:generate stringer -type CastingCMD --output zz_casting_cmd_string.go
type CastingCMD uint8

const (
	CastingReact CastingCMD = 1 + iota
)
