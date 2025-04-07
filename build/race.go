package build

type Race struct {
	Kind RaceCMD
}

//go:generate stringer -type RaceCMD --output zz_race_cmd_string.go
type RaceCMD uint8

const (
	RaceName RaceCMD = 1 + iota
)
