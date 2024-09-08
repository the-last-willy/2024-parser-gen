package tree

type TraverserCommand int

const (
	TraverserContinue TraverserCommand = iota
	TraverserSkip
	TraverserStop
)

// TODO Should iterate edges instead
// Enter, Exit, how to handle root ?
type Traverser[Data any] struct {
	Enter  func(Tree[Data]) TraverserCommand
	Branch func(Tree[Data], int) TraverserCommand
	Exit   func(Tree[Data]) TraverserCommand
}

func NewTraverser[Data any]() *Traverser[Data] {
	return &Traverser[Data]{
		Enter: func(s Tree[Data]) TraverserCommand {
			return TraverserContinue
		},
		Branch: func(Tree[Data], int) TraverserCommand {
			return TraverserContinue
		},
		Exit: func(Tree[Data]) TraverserCommand {
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
