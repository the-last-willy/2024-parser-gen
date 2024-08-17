package mckeeman

import "fmt"

type Name struct {
	Value string
}

func NewName(s string) Name {
	return Name{s}
}

func (n Name) Accept(g Grammar, s string) int {
	r := g.RuleWithName(n)
	return r.Accept(g, s)
}

func (n Name) Print() {
	fmt.Print(n.Value)
}
