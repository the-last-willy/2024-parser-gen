package tree

import "fmt"

type GoFormatter[Data any] struct {
	FormatData func(Data) string
}

func (f GoFormatter[Data]) Format(t Tree[Data]) string {
	return f.format(NewSubTree(t, t.Root()), 0)
}

func (f GoFormatter[Data]) format(st SubTree[Data], indent int) string {
	s := fmt.Sprintf(
		`st.NewNode(
%s,
    []tree.Node{
%s
    },
)`,
		f.formatData(st),
		f.formatChildren(st),
	)
	return IndentStr(s, indent)
}

func (f GoFormatter[Data]) formatData(st SubTree[Data]) string {
	r := st.Root()
	d := st.DataOf(r)
	s := f.FormatData(d)
	return IndentStr(s, 1)
}

func (f GoFormatter[Data]) formatChildren(st SubTree[Data]) string {
	s := ""
	for i, c := range st.ChildrenOf(st.Root()) {
		if i != 0 {
			s += "\n"
		}
		s += f.format(st.WithRoot(c), 2) + ","
	}
	return s
}
