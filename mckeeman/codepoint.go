package mckeeman

import (
	"fmt"
	"strings"
)

type CodePoint struct {
	Rune rune
}

func (c CodePoint) Accept(s string) int {
	if len(s) == 0 {
		return -1
	}
	if rune(s[0]) == c.Rune {
		return 1
	}
	return -1
}

func (c CodePoint) Print() {
	switch c.Rune {
	case ' ':
		fmt.Printf(`0020`)
	case '\n':
		fmt.Printf(`000A`)
	case '\r':
		fmt.Printf(`000D`)
	case '\t':
		fmt.Printf(`0009`)
	default:
		if c.Rune <= 127 {
			fmt.Printf(`%c`, c.Rune)
		} else {
			fmt.Print(strings.ToUpper(fmt.Sprintf(`%x`, c.Rune)))
		}

	}
}
