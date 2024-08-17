package mckeeman

import "fmt"

type Exclude struct {
	Literal Literal

	Next *Exclude
}

func (e Exclude) Accept(g Grammar, s string) int {
	if len(s) == 0 {
		return 0
	}
	if e.Next != nil && e.Next.Accept(g, s) == -1 {
		return -1
	}
	if e.Literal.Accept(g, s) == -1 {
		return 1
	}
	return -1
}

func (e Exclude) Print() {
	fmt.Print(" - ")
	e.Literal.Print()
	if e.Next != nil {
		e.Next.Print()
	}
}
