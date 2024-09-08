package parse

import "parsium/tree"

type Derecursifier struct {
}

func NewDerecursifier() *Derecursifier {
	return &Derecursifier{}
}

func (d *Derecursifier) Process(t tree.Tree[TreeData], b tree.Builder[TreeData]) tree.Builder[TreeData] {
	return b.Tree(
		t.DataOf(*t.Root()),
		d.process(tree.NewSubTree(t, t.Root()), b),
	)
}

// process returns the derecursified children.
func (d *Derecursifier) process(parent tree.SubTree[TreeData], b tree.Builder[TreeData]) []tree.Builder[TreeData] {
	children := []tree.Builder[TreeData]{}
	for _, childNode := range tree.RootChildren[TreeData](parent) {
		child := parent.WithRoot(&childNode)
		db := d.process(child, b)
		if TypeOf(child) == TypeOf(parent) {
			children = append(children, db...)
		} else {
			children = append(
				children,
				b.Tree(
					DataOf(child),
					db,
				),
			)
		}
	}
	return children
}
