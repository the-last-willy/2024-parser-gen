package mckeeman

import "fmt"

type Rule struct {
	name         string
	Alternatives Alternatives
}

func NewRule(name Name, alternatives []Alternative) Rule {
	return Rule{
		name:         name.Value,
		Alternatives: Alternatives{Alternatives: alternatives},
	}
}

func (r Rule) Accept(g Grammar, s string) int {
	return r.Alternatives.Accept(g, s)
}

func (r Rule) Name() Name {
	return NewName(r.name)
}

func (r Rule) Print() {
	r.Name().Print()
	fmt.Print("\n")
	r.Alternatives.Print()
}
