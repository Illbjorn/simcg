package tokenize

/*------------------------------------------------------------------------------
 * Token Set
 *----------------------------------------------------------------------------*/

func NewSet(initialSize int) *Set {
	if initialSize <= 0 {
		initialSize = 64
	}

	return &Set{
		Pos:    -1,
		Tokens: make([]Token, 0, initialSize),
	}
}

type Set struct {
	Pos    int
	Tokens []Token
}

func (set *Set) Adv() {
	next := set.Peek(1)
	if next.Kind == 0 {
		return
	}
	set.Pos += 1
}

var TokenEOF = Token{
	Kind: 0,
}

func (set *Set) Peek(i int) Token {
	i += set.Pos
	if i < 0 {
		return TokenEOF
	} else if i >= len(set.Tokens) {
		return TokenEOF
	} else {
		return set.Tokens[i]
	}
}

func (set *Set) Push(tk Token) {
	set.Tokens = append(set.Tokens, tk)
}

type Snapshot = int

func (set *Set) Snap() Snapshot {
	return set.Pos
}

func (set *Set) Revert(snap Snapshot) {
	set.Pos = snap
}
