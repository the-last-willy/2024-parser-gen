package tree

// Node type is implementation defined.
type Node struct {
	Impl any
}

type Tree[Data any] interface {
	Root() *Node

	ChildrenOf(Node) []Node
	DataOf(Node) Data
}

// Helper functions

func RootChildren[Data any](t Tree[Data]) []Node {
	return t.ChildrenOf(*t.Root())
}

func RootData[Data any](t Tree[Data]) Data {
	return t.DataOf(*t.Root())
}
