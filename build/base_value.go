package build

type BaseValue struct {
	Value string
	Kind  BaseValueKind
}

func (b *BaseValue) String() string {
	return b.Value
}

//go:generate stringer -type BaseValueKind --output zz_base_value_kind_string.go
type BaseValueKind uint8

const (
	BaseID BaseValueKind = 1 + iota
	BaseNum
)
