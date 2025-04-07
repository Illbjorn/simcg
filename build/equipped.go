package build

type Equipped struct {
	ID   string
	Kind EquippedCMD
}

//go:generate stringer -type EquippedCMD --output zz_equipped_cmd_string.go
type EquippedCMD uint8

const (
	EquippedID EquippedCMD = 1 + iota
	EquippedName
	EquippedOnEquipEffect
)
