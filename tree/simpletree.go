package tree

type SimpleTree[Data any] struct {
	root *simpleTreeNode[Data]
}

type simpleTreeNode[Data any] struct {
	data     Data
	children []Node
}

func NewSimpleTree[Data any]() SimpleTree[Data] {
	return SimpleTree[Data]{}
}

func (t SimpleTree[Data]) AsTree() Tree[Data] {
	return NewSimpleTree[Data]()
}

func (t SimpleTree[Data]) Root() Node {
	return Node{
		Impl: t.root,
	}
}

func (t SimpleTree[Data]) IsEmpty() bool {
	return t.root == nil
}

func (t SimpleTree[Data]) ChildrenOf(n Node) []Node {
	return n.Impl.(*simpleTreeNode[Data]).children
}

func (t SimpleTree[Data]) DataOf(n Node) Data {
	return n.Impl.(*simpleTreeNode[Data]).data
}

func (t SimpleTree[Data]) NewNode(data Data, children []Node) Node {
	return Node{
		Impl: &simpleTreeNode[Data]{
			data:     data,
			children: children,
		},
	}
}

func (t SimpleTree[Data]) WithRoot(n Node) Tree[Data] {
	return &SimpleTree[Data]{
		root: n.Impl.(*simpleTreeNode[Data]),
	}
}

func (t SimpleTree[Data]) New(data Data, children []Node) Tree[Data] {
	return t.WithRoot(t.NewNode(data, children))
}
