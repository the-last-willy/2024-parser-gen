package parse

import (
	"parsium/stack"
)

type BottomUpParser struct {
	Grammar func(stack stack.Stack[Token], done bool) *TreeData

	Output func(TreeData)

	Stack stack.Stack[Token]
}

func NewBottomUpParser() *BottomUpParser {
	return &BottomUpParser{
		Stack: stack.NewStack("", "", ""),
	}
}

func (p *BottomUpParser) Process(token Token) {
	p.Stack = p.Stack.Push(token)
	for p.reduce(false) {
	}
}

func (p *BottomUpParser) Flush() {
	for p.reduce(true) {
	}
}

func (p *BottomUpParser) reduce(done bool) bool {
	d := p.Grammar(p.Stack, done)
	if d == nil {
		return false
	}

	cp := *d
	cp.First = p.Stack.Len() - d.Last
	cp.Last = p.Stack.Len() - d.First

	cp.First -= 3
	cp.Last -= 3

	prefix := stack.Slice(p.Stack, nil, d.First)
	suffix := stack.Slice(p.Stack, d.Last, nil)
	p.Stack = stack.NewStack(append(append(prefix, d.Type), suffix...)...)

	p.Output(cp)

	return true
}
