package parse

import "parsium/tree"

type Derecursifier struct {
}

func NewDerecursifier() *Derecursifier {
	return &Derecursifier{}
}

func (d *Derecursifier) Process(t tree.Tree[TreeData]) tree.Tree[TreeData] {
	parent := t.Root()
	children := []tree.Node{}
	for _, child := range t.ChildrenOf(parent) {
		dc := d.Process(t.WithRoot(child))
		if TypeOf(dc) == TypeOf(t) {
			children = append(children, tree.RootChildren(dc)...)
		} else {
			children = append(children, dc.Root())
		}
	}
	return t.WithRoot(t.NewNode(t.DataOf(parent), children))
}
