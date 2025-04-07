package tokenize

import (
	"github.com/illbjorn/echo"
)

var TokenEmpty = Token{}

/*------------------------------------------------------------------------------
 * Tokenize
 *----------------------------------------------------------------------------*/

func Tokenize(data []byte) *Set {
	t := NewTokenizer(data)
	set := NewSet(0)

	for {
		if next := t.Peek(1); next == ByteEOF {
			return set
		}

		// Look for characters to discard first
		if ok := TryDiscard(t); ok {
			continue
		}

		if tk, ok := TryWord(t); ok {
			set.Push(tk)
		} else if tk, ok = TryNum(t); ok {
			set.Push(tk)
		} else if tk, ok = TrySymbol(t); ok {
			set.Push(tk)
		} else if ok = TryLineComment(t); ok {
			continue
		} else {
			panic("impossible")
		}
	}
}

/*------------------------------------------------------------------------------
 * Words
 *----------------------------------------------------------------------------*/

func TryWord(t *Tokenizer) (Token, bool) {
	if !isAlpha(t.Peek(1)) {
		return TokenEmpty, false
	}

	// Mark out locations
	var (
		sy    = t.Line
		sx    = t.Col
		snap  = t.Snapshot()
		start = t.Pos + 1
		stop  = start
	)

	// Consume the word
	for isAlphaNum(t.Peek(1)) {
		t.Pos += 1
		stop += 1
	}
	if stop <= start {
		echo.Error("Consumed zero bytes in TryWord call.")
		t.Revert(snap)
		return TokenEmpty, false
	}

	// Slice off the word, and attempt to resolve it's kind
	word := t.Data[start:stop]
	kind, ok := ResolveKind(word)
	if !ok {
		kind = ID
	}

	// Form the token and return
	return Token{
		Sy:       sy,
		Sx:       sx,
		Ey:       t.Line,
		Ex:       t.Col,
		PosStart: start,
		PosStop:  stop,
		Kind:     kind,
		Source:   t.Data,
	}, true
}

/*------------------------------------------------------------------------------
 * Numbers
 *----------------------------------------------------------------------------*/

func TryNum(t *Tokenizer) (Token, bool) {
	if !isNum(t.Peek(1)) {
		return TokenEmpty, false
	}

	// Mark out locations
	var (
		sy    = t.Line
		sx    = t.Col
		snap  = t.Snapshot()
		start = t.Pos + 1
		stop  = t.Pos + 1
	)

	// Consume the number
	var (
		shouldBreak  bool
		shouldReturn bool
		foundDot     bool
	)
	for {
		shouldBreak, shouldReturn, foundDot = tryNumEval(t, foundDot)
		if shouldBreak {
			break
		}
		if shouldReturn {
			t.Revert(snap)
			return TokenEmpty, false
		}
		t.Pos += 1
		stop += 1
	}

	// Form and return the token
	return Token{
		Kind:     Num,
		Source:   t.Data,
		PosStart: start,
		PosStop:  stop,
		Sx:       sx,
		Sy:       sy,
		Ey:       t.Line,
		Ex:       t.Col,
	}, true
}

var tryNumEval = func(t *Tokenizer, foundDot bool) (
	shouldBreak bool,
	shouldReturn bool,
	foundD bool,
) {
	next := t.Peek(1)

	// End of number
	if !(isNum(next) || next == '.') {
		return true, false, foundDot
	}

	// End of number
	if next == '.' {
		// We already have a float captured
		if foundDot {
			return true, false, foundDot
		}

		// Periods must be immediately followed by another number
		foundDot = true
		nextNext := t.Peek(2)
		if !isNum(nextNext) {
			return true, false, foundDot
		}
	}

	// Ok
	return false, false, foundDot
}

/*------------------------------------------------------------------------------
 * Symbols
 *----------------------------------------------------------------------------*/

func TrySymbol(t *Tokenizer) (Token, bool) {
	if !isSymbol(t.Peek(1)) {
		return TokenEmpty, false
	}

	// Mark out locations
	var (
		sy    = t.Line
		sx    = t.Col
		snap  = t.Snapshot()
		start = t.Pos + 1
	)

	// Get the next 3-bytes (maximum possible symbol length)
	var (
		next1 = t.Peek(1)
		next2 = t.Peek(2)
		next3 = t.Peek(3)
		data  = []byte{next1, next2, next3}

		kind Kind
		ok   bool
	)

	// Attempt 3-byte symbol match
	if kind, ok = ResolveKind(data); !ok {
		// Attempt 2-byte symbol match
		data = data[:len(data)-1]
		if kind, ok = ResolveKind(data); !ok {
			// Attempt 1-byte symbol match
			data = data[:len(data)-1]
			if kind, ok = ResolveKind(data); !ok {
				t.Revert(snap)
				return TokenEmpty, false
			}
		}
	}

	// Advance tokenizer
	for range len(data) {
		t.Adv()
	}

	// Form and return the token
	return Token{
		Kind:     kind,
		Source:   t.Data,
		PosStart: start,
		PosStop:  t.Pos + 1,
		Sx:       sx,
		Sy:       sy,
		Ey:       t.Line,
		Ex:       t.Col,
	}, true
}

/*------------------------------------------------------------------------------
 * Line Comments
 *----------------------------------------------------------------------------*/

func TryLineComment(t *Tokenizer) bool {
	if t.Peek(1) != '#' {
		return false
	}

	// Consume til' the end-of-line
	for t.Peek(1) != '\n' {
		t.Adv()
	}

	return true
}

/*------------------------------------------------------------------------------
 * Discards
 *----------------------------------------------------------------------------*/

var discards = []byte("., \t\r\n\f")

var isDiscard = func(b byte) bool {
	for i := range discards {
		if b == discards[i] {
			return true
		}
	}
	return false
}

func TryDiscard(t *Tokenizer) bool {
	if !isDiscard(t.Peek(1)) {
		return false
	}

	for isDiscard(t.Peek(1)) {
		t.Adv()
	}

	return true
}

/*------------------------------------------------------------------------------
 * Whitespace
 *----------------------------------------------------------------------------*/

func TrySpace(t *Tokenizer) bool {
	if !isSpace(t.Peek(1)) {
		return false
	}

	for isSpace(t.Peek(1)) {
		t.Adv()
	}

	return true
}

/*------------------------------------------------------------------------------
 * Byte Recognition
 *----------------------------------------------------------------------------*/

func isAlpha(b byte) bool {
	return b >= 0x41 && b <= 0x5a || b >= 0x61 && b <= 0x7a
}

func isNum(b byte) bool {
	return b >= 0x30 && b <= 0x39
}

func isAlphaNum(b byte) bool {
	return isAlpha(b) || isNum(b) || b == '_'
}

const (
	Tab            = '\t'
	Space          = ' '
	Newline        = '\n'
	CarriageReturn = '\r'
	FormFeed       = '\f'
)

func isSpace(b byte) bool {
	return b == Tab || b == Space || b == Newline || b == CarriageReturn || b == FormFeed
}

func isSymbol(b byte) bool {
	return !(isAlphaNum(b) || isSpace(b))
}
