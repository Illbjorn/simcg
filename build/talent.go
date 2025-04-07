package build

type Talent struct {
	ID   string
	Kind TalentCMD
}

//go:generate stringer -type TalentCMD --output zz_talent_cmd_string.go
type TalentCMD uint8

const (
	TalentEnabled TalentCMD = iota
	TalentRank
)
