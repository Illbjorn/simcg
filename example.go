package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/illbjorn/echo"

	"github.com/illbjorn/simcg/build"
	"github.com/illbjorn/simcg/parse"
	"github.com/illbjorn/simcg/tokenize"
)

func main() {
	echo.SetFlags(echo.WITH_CALLER_FILE, echo.WITH_CALLER_LINE)
	echo.SetLevel(echo.LevelWarn)

	args := ParseArgs()

	content := read(args.Input)
	set := tokenize.Tokenize(content)
	nodes := parse.Parse(set)
	list := build.Build(nodes)
	for _, instruction := range list.Instructions {
		if instruction.Executor != nil {
			echo.Infof(">> %s", instruction.Executor.ID)
			printCommand(instruction.Executor.Command, 0)
			fmt.Println()
		}
	}

	echo.Info("Done!")
}

func printCommand(c *build.Command, i int) {
	if c == nil {
		return
	}
	pad := "  - "
	switch c.Kind {
	case 0:
		echo.Fatalf(pad+"Found command with no kind ['%#v'].", c)
	case build.CMDTypeBuiltin:
		echo.Infof(pad+"%s", c.Builtin)
	case build.CMDTypeGCD:
		echo.Infof(pad+"%s", c.GCD)
	case build.CMDTypeTarget:
		echo.Infof(pad+"%s", c.Target)
	case build.CMDTypeCooldown:
		echo.Infof(pad+"%s", c.Cooldown)
	case build.CMDTypeBuff:
		echo.Infof(pad+"%s", c.Buff)
	case build.CMDTypeDebuff:
		echo.Infof(pad+"%s", c.Debuff)
	case build.CMDTypeOp:
		echo.Infof(pad+"%s", c.Op)
	case build.CMDTypeBase:
		echo.Infof(pad+"%s", c.Base)
	case build.CMDTypeMovement:
		echo.Infof(pad+"%s", c.Movement)
	case build.CMDTypeVariable:
		echo.Infof(pad+"%s", c.Variable)
	case build.CMDTypeDot:
		echo.Infof(pad+"%s", c.Dot)
	case build.CMDTypeTalent:
		echo.Infof(pad+"%s", c.Talent)
	case build.CMDTypeRaidEvent:
		echo.Infof(pad+"%s", c.RaidEvent)
	default:
		echo.Fatalf(pad+"Found unexpected command kind ['%s'].", c.Kind)
	}

	printCommand(c.Next, i+1)
}

type Args struct {
	Input string
}

func ParseArgs() Args {
	var args Args

	// --input, -i
	flag.StringVar(&args.Input, "input", "", "")
	flag.StringVar(&args.Input, "i", "", "")
	flag.Parse()

	return args
}

func read(path string) []byte {
	rc := openRead(path)
	defer rc.Close()

	if data, err := io.ReadAll(rc); err != nil {
		echo.Fatalf("Failed to read input file ['%s']: %s.", path, err)
		panic("impossible")
	} else {
		return data
	}
}

func openRead(path string) io.ReadCloser {
	f, err := os.Open(path)
	if err != nil {
		echo.Fatalf("Failed to open input file ['%s']: %s.", path, err)
	}

	return f
}
