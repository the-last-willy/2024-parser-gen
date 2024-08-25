package tree

import "fmt"

type simpleNode[Data any] struct {
	data     Data
	children []simpleNode[Data]
}

// SimpleTree
// Node.Impl type is SimpleTree[Data].
type SimpleTree[Data any] struct {
	root *simpleNode[Data]
}

func NewSimpleTree[Data any]() SimpleTree[Data] {
	return SimpleTree[Data]{}
}

func (t SimpleTree[Data]) AsTree() Tree[Data] {
	return t
}

func (t SimpleTree[Data]) Root() Node {
	return Node{
		Impl: t.root,
	}
}

func (t SimpleTree[Data]) IsEmpty() bool {
	return false
}

func (t SimpleTree[Data]) ChildrenOf(n Node) []Node {
	children := []Node{}
	fmt.Printf("%p\n", n.Impl)
	for i, _ := range n.Impl.(*simpleNode[Data]).children {
		c := &n.Impl.(*simpleNode[Data]).children[i]
		fmt.Printf("%p\n", c)
		children = append(children, Node{Impl: c})
	}
	return children
}

func (t SimpleTree[Data]) DataOf(n Node) Data {
	return n.Impl.(*simpleNode[Data]).data
}

func (t SimpleTree[Data]) NewNode(data Data, children []Node) Node {
	realChildren := []simpleNode[Data]{}
	for _, c := range children {
		realChildren = append(realChildren, *c.Impl.(*simpleNode[Data]))
	}
	return Node{
		Impl: &simpleNode[Data]{
			data:     data,
			children: realChildren,
		},
	}
}

func (t SimpleTree[Data]) WithRoot(n Node) Tree[Data] {
	return &SimpleTree[Data]{
		root: n.Impl.(*simpleNode[Data]),
	}
}

func (t SimpleTree[Data]) New(data Data, children []Node) Tree[Data] {
	return t.WithRoot(t.NewNode(data, children))
}
