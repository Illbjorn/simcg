package build

type Target struct {
	Kind   TargetCMD
	Filter TargetFilterCMD
	Debuff *Debuff
	Buff   *Buff
}

//go:generate stringer -type TargetCMD --output zz_target_cmd_string.go
type TargetCMD uint8

const (
	TargetLevel TargetCMD = 1 + iota
	TargetHealth
	TargetAdds
	TargetAddsNever
	TargetDistance
	TargetCurrentTarget
	TargetName
	TargetTTD
)

//go:generate stringer -type TargetFilterCMD --output zz_target_filter_cmd_string.go
type TargetFilterCMD uint8

const (
	TargetFilterPct TargetFilterCMD = 1 + iota
)
