package build

type Spec struct {
	Kind SpecCMD
}

//go:generate stringer -type SpecCMD --output zz_spec_cmd_string.go
type SpecCMD uint8

const (
	SpecName SpecCMD = 1 + iota
)
