package mckeeman

import "fmt"

type Alternative struct {
	Items Items
}

func NewAlternative(items []Item) Alternative {
	return Alternative{Items: Items{Items: items}}
}

func EmptyAlternative() Alternative {
	return NewAlternative([]Item{
		NewCharacters(""),
	})
}

func (a Alternative) Accept(g Grammar, s string) int {
	return a.Items.Accept(g, s)
}

func (a Alternative) Print() {
	fmt.Print("    ")
	a.Items.Print()
	fmt.Print("\n")
}
