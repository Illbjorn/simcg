package build

type Role struct {
	Kind RoleCMD
}

//go:generate stringer -type RoleCMD --output zz_role_cmd_string.go
type RoleCMD uint8

const (
	RoleAttack RoleCMD = 1 + iota
	RoleHeal
	RoleSpell
	RoleTank
)
