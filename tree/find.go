package tree

func FindAll[Data any](t Tree[Data], pred func(Tree[Data]) bool) []Node {
	found := []Node{}
	if pred(t) {
		found = append(found, t.Root())
	}
	for _, c := range t.ChildrenOf(t.Root()) {
		found = append(found, FindAll(t.WithRoot(c), pred)...)
	}
	return found
}

func FindFirst[Data any](t Tree[Data], pred func(Tree[Data]) bool) *Node {
	if pred(t) {
		r := t.Root()
		return &r
	}
	for _, c := range t.ChildrenOf(t.Root()) {
		f := FindFirst(t.WithRoot(c), pred)
		if f != nil {
			return f
		}
	}
	return nil
}
