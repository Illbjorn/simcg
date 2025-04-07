package build

type VarDefinition struct {
	Name      string
	Value     *Command
	ValueElse *Command
}

type VarReference struct {
	ID string
}
