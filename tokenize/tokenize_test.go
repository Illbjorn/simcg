package tokenize

import (
	"os"
	"testing"

	"github.com/illbjorn/echo"
	ass "github.com/stretchr/testify/assert"
)

func init() {
	echo.SetFlags(
		echo.WithCallerFile,
		echo.WithCallerLine,
		echo.WithLevel,
		echo.WithColor,
	)
}

var assert = ass.Equal

func TestTokenize(t *testing.T) {
	data, err := os.ReadFile("sample.simc")
	assert(t, nil, err)

	set := Tokenize(data)
	_ = set
}

func TestTryWord(t *testing.T) {
	tk, ok := TryWord(NewTokenizer([]byte("actions! \n!")))
	assert(t, true, ok)
	assert(t, []byte("actions"), tk.Value())
	assert(t, Actions, tk.Kind)
}

func TestTryNum(t *testing.T) {
	tk, ok := TryNum(NewTokenizer([]byte("1.2")))
	assert(t, true, ok)
	assert(t, []byte("1.2"), tk.Value())
	assert(t, Num, tk.Kind)

	tk, ok = TryNum(NewTokenizer([]byte("1.2.3")))
	assert(t, true, ok)
	assert(t, []byte("1.2"), tk.Value())
}

func TestTrySymbol(t *testing.T) {
	tk, ok := TrySymbol(NewTokenizer([]byte("+=/")))
	assert(t, true, ok)
	assert(t, []byte("+=/"), tk.Value())
	assert(t, AddAssign, tk.Kind)
}
