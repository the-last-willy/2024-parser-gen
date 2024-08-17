package parse

type DataTree struct {
	Root  *Tree
	Store string
}

func (dt DataTree) DataOf(t *Tree) string {
	return dt.Store[t.First:t.Last]
}

func (g DataTree) WithRoot(t *Tree) DataTree {
	return DataTree{
		Root:  t,
		Store: g.Store,
	}
}
