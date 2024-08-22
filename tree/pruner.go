package tree

type Pruner[Data any] struct {
	Remove             func(Tree[Data], Node) bool
	RemoveChildren     func(Tree[Data], Node) bool
	RemoveKeepChildren func(Tree[Data], Node) bool
}

func (p *Pruner[Data]) Apply(tree Tree[Data]) Tree[Data] {
	if tree.IsEmpty() {
		return tree
	}

	pruned := p.apply(tree, tree.Root())

	if len(pruned) != 1 {
		panic("Pruner.Apply: bad pruning")
	}

	return tree.WithRoot(pruned[0])
}

func (p *Pruner[Data]) apply(t Tree[Data], n Node) []Node {
	if p.Remove(t, n) {
		return []Node{}
	}

	if p.RemoveChildren(t, n) {
		newN := t.NewNode(t.DataOf(n), []Node{})
		return []Node{newN}
	}

	pruned := []Node{}
	for _, child := range t.ChildrenOf(n) {
		pruned = append(pruned, p.apply(t, child)...)
	}

	if p.RemoveKeepChildren(t, n) {
		return pruned
	}

	return []Node{t.NewNode(t.DataOf(n), pruned)}
}
