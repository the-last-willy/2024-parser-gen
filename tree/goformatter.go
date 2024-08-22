package tree

import "fmt"

type GoFormatter[Data any] struct {
	FormatData func(Data) string
}

func (f GoFormatter[Data]) Format(t Tree[Data]) string {
	return f.format(t, 0)
}

func (f GoFormatter[Data]) format(t Tree[Data], indent int) string {
	s := fmt.Sprintf(
		`t.NewNode(
%s,
    []tree.Node{
%s
    },
)`,
		f.formatData(t, indent),
		f.formatChildren(t, indent),
	)
	return IndentStr(s, indent)
}

func (f GoFormatter[Data]) formatData(t Tree[Data], indent int) string {
	r := t.Root()
	d := t.DataOf(r)
	s := f.FormatData(d)
	return IndentStr(s, 1)
}

func (f GoFormatter[Data]) formatChildren(t Tree[Data], indent int) string {
	s := ""
	for i, c := range t.ChildrenOf(t.Root()) {
		if i != 0 {
			s += "\n"
		}
		s += f.format(t.WithRoot(c), 2) + ","
	}
	return s
}
