package tree

// Node type is implementation defined.
type Node struct {
	Impl any
}

type Tree[Data any] interface {
	Root() Node

	IsEmpty() bool

	ChildrenOf(Node) []Node
	DataOf(Node) Data

	// TODO Remove, this one can't be functional
	NewNode(data Data, children []Node) Node

	// TODO Move to a tree builder ???
	WithRoot(Node) Tree[Data]

	// TODO Move to a tree builder
	New(data Data, children []Node) Tree[Data]
}

func RootChildren[Data any](t Tree[Data]) []Node {
	return t.ChildrenOf(t.Root())
}

func RootData[Data any](t Tree[Data]) Data {
	return t.DataOf(t.Root())
}
