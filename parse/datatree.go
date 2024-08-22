package parse

type DataTree struct {
	Root  *Tree
	Store string
}

func (dt DataTree) ChildrenOf(n any) []any {
	children := []any{}
	for _, child := range n.(*Tree).Children {
		children = append(children, child)
	}
	return children
}

func (dt DataTree) DataOf(n any) string {
	t := n.(*Tree)
	return dt.Store[t.First:t.Last]
}

func (dt DataTree) IsEmpty() bool {
	return dt.Root == nil
}

func (g DataTree) WithRoot(t *Tree) DataTree {
	return DataTree{
		Root:  t,
		Store: g.Store,
	}
}

func (g DataTree) NewNode(data Tree, children []*Tree) *Tree {
	data.Children = children
	return &data
}

func (g DataTree) GetRoot() any {
	return g.Root
}
