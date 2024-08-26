package parse

import "parsium/tree"

type Derecursifier struct {
}

func NewDerecursifier() *Derecursifier {
	return &Derecursifier{}
}

func (d *Derecursifier) Process(t tree.Tree[TreeData], b tree.Builder[TreeData]) tree.Builder[TreeData] {
	return b.Tree(
		t.DataOf(t.Root()),
		d.process(t, t.Root(), b),
	)
}

// process returns the derecursified children.
func (d *Derecursifier) process(t tree.Tree[TreeData], n tree.Node, b tree.Builder[TreeData]) []tree.Builder[TreeData] {
	children := []tree.Builder[TreeData]{}
	for _, child := range t.ChildrenOf(n) {
		db := d.process(t, child, b)
		if TypeOf(t, child) == TypeOf(t, n) {
			children = append(children, db...)
		} else {
			children = append(
				children,
				b.Tree(
					t.DataOf(child),
					db,
				),
			)
		}
	}
	return children
}
