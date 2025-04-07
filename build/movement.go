package build

type Movement struct {
	Kind MovementCMD
}

//go:generate stringer -type MovementCMD --output zz_movement_cmd_string.go
type MovementCMD uint8

const (
	MovementRemains MovementCMD = 1 + iota
	MovementDistance
	MovementSpeed
)
