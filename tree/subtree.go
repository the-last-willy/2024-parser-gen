package tree

type SubTree[Data any] struct {
	tree Tree[Data]
	root Node
}

func NewSubTree[Data any](t Tree[Data], root Node) SubTree[Data] {
	return SubTree[Data]{
		tree: t,
		root: root,
	}
}

func (st SubTree[Data]) WithRoot(root Node) SubTree[Data] {
	return NewSubTree(st.tree, root)
}

// Tree interface

func (st SubTree[Data]) ChildrenOf(n Node) []Node {
	return st.tree.ChildrenOf(n)
}

func (st SubTree[Data]) DataOf(n Node) Data {
	return st.tree.DataOf(n)
}

func (st SubTree[Data]) IsEmpty() bool {
	return st.tree.IsEmpty()
}

func (st SubTree[Data]) Root() Node {
	return st.root
}
