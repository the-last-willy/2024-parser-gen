package mckeeman

import (
	"fmt"
	"strings"
)

type Characters struct {
	CodePoints []CodePoint
	Runes      []rune
}

func NewCharacters(s string) Characters {
	cs := make([]CodePoint, len(s))
	for i, r := range s {
		cs[i] = CodePoint{Rune: r}
	}
	return Characters{
		CodePoints: cs,
		Runes:      []rune(s),
	}
}

func (c Characters) Accept(_ Grammar, s string) int {
	if strings.HasPrefix(s, string(c.Runes)) {
		return len(c.Runes)
	}
	return -1
}

func (c Characters) Print() {
	fmt.Print(`"`)
	for _, r := range c.Runes {
		fmt.Printf("%c", r)
	}
	fmt.Print(`"`)
}
