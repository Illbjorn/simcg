package build

type Op uint8

const (
	// Equality
	OpGE Op = 1 + iota
	OpLE
	OpGT
	OpLT
	OpEQ
	OpNE

	// Logical
	OpNot
	OpAnd
	OpOr

	// Arithmetic
	OpMod
	OpMult
	OpDiv
	OpAdd
	OpSub

	// Misc
	OpMax
	OpMin
)

func (o Op) String() string {
	switch o {
	// Equality
	case OpGE:
		return ">="
	case OpLE:
		return "<="
	case OpGT:
		return ">"
	case OpLT:
		return "<"
	case OpEQ:
		return "="
	case OpNE:
		return "!="

	// Logical
	case OpNot:
		return "!"
	case OpAnd:
		return "&"
	case OpOr:
		return "|"

	// Arithmetic
	case OpMod:
		return "%%"
	case OpMult:
		return "*"
	case OpDiv:
		return "%"
	case OpAdd:
		return "+"
	case OpSub:
		return "-"

	// Misc
	case OpMax:
		return ">?"
	case OpMin:
		return "<?"
	default:
		panic("Ok")
	}
}
