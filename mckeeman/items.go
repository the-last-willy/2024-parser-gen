package mckeeman

import "fmt"

type Items struct {
	Items []Item
}

func (i Items) Accept(g Grammar, s string) int {
	offset := 0
	for _, i := range i.Items {
		c := i.Accept(g, s[offset:])
		if c == -1 {
			return -1
		}
		offset += c
	}
	return offset
}

func (i Items) Print() {
	for idx, i := range i.Items {
		if idx > 0 {
			fmt.Printf(" ")
		}
		i.Print()
	}
}
