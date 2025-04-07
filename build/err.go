package build

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/illbjorn/echo"

	"github.com/illbjorn/simcg/parse"
)

/*------------------------------------------------------------------------------
 * Errors
 *----------------------------------------------------------------------------*/

func kindErr(n *parse.Node) {
	file, line, fn := caller(2)
	echo.Fatalf(
		"['%s::%s/%s'] Found unexpected node kind ['%s']:\n%s.",
		file, line, fn, n.Kind, n,
	)
	os.Exit(1)
}

func caller(skip int) (string, string, string) {
	pcs := make([]uintptr, 1)
	runtime.Callers(skip+1, pcs)
	frames := runtime.CallersFrames(pcs)
	frame, _ := frames.Next()

	i := strings.LastIndexByte(frame.File, '/')
	file := frame.File[i+1:]

	line := frame.Line

	fn := frame.Func.Name()

	return file, strconv.Itoa(line), fn
}
