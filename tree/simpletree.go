package tree

// SimpleTree
// Node.Impl type is SimpleTree[Data].
type SimpleTree[Data any] struct {
	data Data

	children []Node
}

func NewSimpleTree[Data any]() SimpleTree[Data] {
	return SimpleTree[Data]{}
}

func newSimpleTree[Data any](d Data, children []SimpleTree[Data]) SimpleTree[Data] {
	nodes := []Node{}
	for _, child := range children {
		nodes = append(nodes, Node{Impl: child})
	}
	return SimpleTree[Data]{
		data:     d,
		children: nodes,
	}
}

func (t SimpleTree[Data]) AsTree() Tree[Data] {
	return t
}

func (t SimpleTree[Data]) Root() Node {
	return Node{
		Impl: t,
	}
}

func (t SimpleTree[Data]) IsEmpty() bool {
	return false
}

func (t SimpleTree[Data]) ChildrenOf(n Node) []Node {
	return n.Impl.(SimpleTree[Data]).children
}

func (t SimpleTree[Data]) DataOf(n Node) Data {
	return n.Impl.(SimpleTree[Data]).data
}

func (t SimpleTree[Data]) NewNode(data Data, children []Node) Node {
	return Node{
		Impl: SimpleTree[Data]{
			data:     data,
			children: children,
		},
	}
}

func (t SimpleTree[Data]) WithRoot(n Node) Tree[Data] {
	return &SimpleTree[Data]{
		data:     t.DataOf(n),
		children: t.ChildrenOf(n),
	}
}

func (t SimpleTree[Data]) New(data Data, children []Node) Tree[Data] {
	return t.WithRoot(t.NewNode(data, children))
}
