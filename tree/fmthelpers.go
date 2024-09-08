package tree

import (
	"strings"
)

func IndentStr(str string, level int) string {
	lines := strings.Split(str, "\n")
	indent := strings.Repeat("    ", level)

	s := ""
	for i, l := range lines {
		if i != 0 {
			s += "\n"
		}
		if l != "" {
			s += indent + l
		}
	}

	return s
}
