package tree

type TraverserCommand int

const (
	TraverserContinue TraverserCommand = iota
	TraverserSkip
	TraverserStop
)

type Traverser[Data any] struct {
	Enter  func(SubTree[Data]) TraverserCommand
	Branch func(SubTree[Data], int) TraverserCommand
	Exit   func(SubTree[Data]) TraverserCommand
}

func NewTraverser[Data any]() *Traverser[Data] {
	return &Traverser[Data]{
		Enter: func(s SubTree[Data]) TraverserCommand {
			return TraverserContinue
		},
		Branch: func(s SubTree[Data], i int) TraverserCommand {
			return TraverserContinue
		},
		Exit: func(s SubTree[Data]) TraverserCommand {
			return TraverserContinue
		},
	}
}

func (t *Traverser[Data]) Process(tr Tree[Data]) {
	if tr.Root() != nil {
		t.process(NewSubTree(tr, tr.Root()))
	}
}

func (t *Traverser[Data]) process(st SubTree[Data]) TraverserCommand {
	cmd := t.Enter(st)
	if cmd == TraverserSkip {
		t.Exit(st)
		return TraverserContinue
	}
	if cmd == TraverserStop {
		t.Exit(st)
		return TraverserStop
	}
	for i, child := range RootChildren(st) {
		t.Branch(st, i)
		cmd := t.process(st.WithRoot(&child))
		if cmd == TraverserSkip {
			break
		}
		if cmd == TraverserStop {
			t.Exit(st)
			return TraverserStop
		}
	}
	t.Exit(st)
	return TraverserContinue
}
