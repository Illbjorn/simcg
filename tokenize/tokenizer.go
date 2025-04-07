package tokenize

/*------------------------------------------------------------------------------
 * Tokenizer
 *----------------------------------------------------------------------------*/

const ByteEOF = '\x00'

func NewTokenizer(data []byte) *Tokenizer {
	// Count lines in source
	lines := 0
	for i := range data {
		if data[i] == '\n' {
			lines += 1
		}
	}

	return &Tokenizer{
		Data:        data,
		Pos:         -1,
		Line:        1,
		Col:         1,
		LineIndices: make([][2]int, lines),
	}
}

type Tokenizer struct {
	Data        []byte
	LineIndices [][2]int
	Pos         int
	Line        int
	Col         int
}

func (t *Tokenizer) Adv() {
	next := t.Peek(1)
	if next == ByteEOF {
		return
	}

	t.Pos += 1
	if next == '\n' {
		if t.Line == 0 {
			t.LineIndices = append(t.LineIndices, [2]int{0, t.Pos - 1})
		} else {
			lastLineStop := t.LineIndices[t.Line-1][1]
			t.LineIndices = append(t.LineIndices, [2]int{lastLineStop + 2, t.Pos - 1})
		}
		t.Line += 1
		t.Col = 1
	} else {
		t.Col += 1
	}
}

func (t *Tokenizer) Peek(i int) byte {
	i = t.Pos + i
	if i < 0 {
		return ByteEOF
	} else if i >= len(t.Data) {
		return ByteEOF
	} else {
		return t.Data[i]
	}
}

func (t *Tokenizer) Snapshot() int {
	return t.Pos
}

func (t *Tokenizer) Revert(snap int) {
	t.Pos = snap
}
