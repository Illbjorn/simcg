package build

type Toggle struct {
	ID    string
	State ToggleState
}

type ToggleState uint8

const (
	ON ToggleState = 1 + iota
	OFF
)
