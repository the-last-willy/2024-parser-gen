package tree

type Pruner[Data any] struct {
	Remove             func(Tree[Data], Node) bool
	RemoveChildren     func(Tree[Data], Node) bool
	RemoveKeepChildren func(Tree[Data], Node) bool
}

func (p *Pruner[Data]) Apply(t Tree[Data], b Builder[Data]) Builder[Data] {
	if t.Root() == nil {
		return b.Empty()
	}

	pruned := p.apply(t, *t.Root(), b)

	if len(pruned) != 1 {
		panic("Pruner.Apply: bad pruning")
	}

	return pruned[0]
}

func (p *Pruner[Data]) apply(t Tree[Data], n Node, b Builder[Data]) []Builder[Data] {
	if p.Remove(t, n) {
		return []Builder[Data]{}
	}

	if p.RemoveChildren(t, n) {
		return []Builder[Data]{
			b.Tree(t.DataOf(n), []Builder[Data]{}),
		}
	}

	pruned := []Builder[Data]{}
	for _, child := range t.ChildrenOf(n) {
		pruned = append(pruned, p.apply(t, child, b)...)
	}

	if p.RemoveKeepChildren(t, n) {
		return pruned
	}

	return []Builder[Data]{
		b.Tree(t.DataOf(n), pruned),
	}
}
