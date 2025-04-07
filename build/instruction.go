package build

type Instruction struct {
	ID             string          `json:"id,omitempty"`
	Kind           InstructionKind `json:"kind,omitempty"`
	Variable       *VarDefinition  `json:"var,omitempty"`
	Toggle         *Toggle         `json:"toggle,omitempty"`
	Executor       *Executor       `json:"executor,omitempty"`
	CallActionList *CallActionList `json:"call,omitempty"`
	RunActionList  *RunActionList  `json:"run,omitempty"`
	UseItem        *UseItem        `json:"use,omitempty"`
}

//go:generate stringer -type InstructionKind --output zz_instruction_kind_string.go
type InstructionKind uint8

const (
	KindSnapshotStats InstructionKind = 1 + iota
	KindAutoAttack
	KindVariable
	KindToggle
	KindExecutor
	KindCallActionList
	KindRunActionList
	KindUseItem
)

type CallActionList struct {
	ID string
}

type RunActionList struct {
	ID string
}

type UseItem struct {
	ID         string
	Conditions any
}
