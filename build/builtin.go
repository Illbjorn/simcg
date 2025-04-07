package build

type Builtin struct {
	Kind     BuiltinCMD
	Resource Resource
}

//go:generate stringer -type BuiltinCMD --output zz_builtin_cmd_string.go
type BuiltinCMD uint8

const (
	BuiltinTime BuiltinCMD = 1 + iota
	BuiltinActiveEnemies
	BuiltinExpectedCombatLength
	BuiltinFightRemains
	BuiltinResource
)
