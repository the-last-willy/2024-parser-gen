package tree

// A SimpleBuilder builds the tree on-the-fly.
type SimpleBuilder[Data any] struct {
	tree SimpleTree[Data]
}

func NewSimpleBuilder[Data any]() SimpleBuilder[Data] {
	return SimpleBuilder[Data]{}
}

func (b SimpleBuilder[Data]) Build() SimpleTree[Data] {
	return b.tree
}

// Builder interface

func (b SimpleBuilder[Data]) ConvertTo(other Builder[Data]) Builder[Data] {
	if _, ok := other.(SimpleBuilder[Data]); ok {
		// It already has the expected type, no need to convert it.
		return b
	}

	panic("not implemented")
}

func (b SimpleBuilder[Data]) For(t Tree[Data]) Builder[Data] {
	if simpleTree, ok := t.(SimpleTree[Data]); ok {
		return SimpleBuilder[Data]{
			tree: simpleTree,
		}
	}

	panic("not implemented")
}

func (b SimpleBuilder[Data]) Tree(d Data, children []Builder[Data]) Builder[Data] {
	subtrees := []simpleNode[Data]{}
	for _, c := range children {
		subt := c.ConvertTo(b).(SimpleBuilder[Data]).Build()
		subtrees = append(subtrees, *subt.root)
	}
	return SimpleBuilder[Data]{
		tree: SimpleTree[Data]{
			root: &simpleNode[Data]{
				data:     d,
				children: subtrees,
			},
		},
	}
}
