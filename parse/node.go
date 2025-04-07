package parse

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/illbjorn/echo"
)

type Node struct {
	Kind           Kind
	Sy, Sx, Ey, Ex int
	PosStart, PosStop int
	Source         []byte
	Parent         *Node
	Children       []*Node
}

const (
	tmpNodeHeader = "   %-20s  %-20s  %-20s  %-20s\n"
	tmplNode      = "%-2d %-20s  %-20s  %-20d  %-20d\n"
)

func (n Node) String() string {
	b := new(strings.Builder)
	b.Grow(256)
	fmt.Fprintf(b, tmpNodeHeader, "KIND", "VALUE", "SX", "SY")
	n.string(b, 0)
	return b.String()
}

func (n Node) string(b *strings.Builder, i int) {
	fmt.Fprintf(b, tmplNode, i, n.Kind, Value(&n), n.Sx, n.Sy)
	for child := range Children(&n) {
		child.string(b, i+1)
	}
}

func Value(n *Node) []byte {
	if len(n.Source) == 0 {
		return nil
	}
	return n.Source[n.PosStart:n.PosStop]
}

func ValueS(n *Node) string {
	if len(n.Source) == 0 {
		return ""
	}
	return btos(n.Source[n.PosStart:n.PosStop])
}

func btos(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return unsafe.String(&b[0], len(b))
}

func Source(n *Node) []byte {
	if len(n.Source) == 0 {
		return nil
	}

	lineStart := 0
	for i := n.PosStart; i > 0; i-- {
		if n.Source[i] == '\n' {
			lineStart = i + 1
			break
		}
	}

	lineStop := len(n.Source)
	for i := n.PosStop; i < len(n.Source); i++ {
		if n.Source[i] == '\n' {
			lineStop = i
			break
		}
	}

	echo.Debug("token.Source =>")
	echo.Debugf("%-2s: %s", "Line", n.Source[lineStart:lineStop])
	echo.Debugf("%-2s: %s", "Source", n.Source)
	echo.Debugf("%-2s: %s", "Line Start", lineStart)
	echo.Debugf("%-2s: %s", "Line Stop", lineStop)

	return n.Source[lineStart:lineStop]
}

func Push(branch *Node, leaves ...*Node) {
	for _, leaf := range leaves {
		if leaf.Kind > 0 {
			branch.Children = append(branch.Children, leaf)
		}
	}
}

func Children(n *Node) func(yield func(*Node) bool) {
	return func(yield func(*Node) bool) {
		for _, child := range n.Children {
			if !yield(child) {
				return
			}
		}
	}
}
