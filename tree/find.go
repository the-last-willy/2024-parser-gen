package tree

func FindAll[Data any](st SubTree[Data], pred func(SubTree[Data]) bool) []Node {
	found := []Node{}
	if pred(st) {
		found = append(found, *st.Root())
	}
	for _, c := range st.ChildrenOf(*st.Root()) {
		found = append(found, FindAll(st.WithRoot(&c), pred)...)
	}
	return found
}

func FindFirst[Data any](st SubTree[Data], pred func(SubTree[Data]) bool) *Node {
	if pred(st) {
		r := *st.Root()
		return &r
	}
	for _, c := range st.ChildrenOf(*st.Root()) {
		f := FindFirst(st.WithRoot(&c), pred)
		if f != nil {
			return f
		}
	}
	return nil
}
