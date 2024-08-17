package mckeeman

import "fmt"

type Rules struct {
	Rules []Rule
}

func (r Rules) Accept(g Grammar, s string) int {
	for _, r := range r.Rules {
		a := r.Accept(g, s)
		if a != -1 {
			return a
		}
	}
	return -1
}

func (r Rules) Print() {
	for i, r := range r.Rules {
		if i > 0 {
			fmt.Printf("\n")
		}
		r.Print()
	}
}
