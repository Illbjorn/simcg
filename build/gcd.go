package build

type GCD struct {
	Kind GCDCMD
}

//go:generate stringer -type GCDCMD --output zz_gcd_cmd_string.go
type GCDCMD uint8

const (
	GCDRef GCDCMD = iota
	GCDRemains
	GCDMax
)
