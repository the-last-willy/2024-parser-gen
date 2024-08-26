package parse

import (
	"fmt"
	"parsium/tree"
)

type TreeFormatter struct {
}

func (f TreeFormatter) Format(tr tree.SubTree[TreeData], src string) string {
	r := *tr.Root()
	rd := tr.DataOf(r)
	rc := tr.ChildrenOf(r)

	txt := ""
	if len(rc) == 0 {
		txt = " = " + src[rd.First:rd.Last]
	}

	children := ""
	for _, child := range rc {
		children += f.Format(tr.WithRoot(&child), src)
	}
	children = tree.IndentStr(children, 1)

	s := fmt.Sprintf(
		`(%d, %d) %s%s
%s`,
		rd.First,
		rd.Last,
		rd.Type,
		txt,
		children,
	)

	return s
}
