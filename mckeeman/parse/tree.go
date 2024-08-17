package mck_parse

import (
	"fmt"
	"parsium/log"
	"strings"
)

type Tree struct {
	First int
	Last  int

	Value string

	Children []Tree
}

func (t Tree) Culled() []Tree {
	// Empty nodes are culled
	if t.First == t.Last {
		return []Tree{}
	}
	if t.Value == "#literal" {
		return []Tree{}
	}

	culledChildren := []Tree{}
	for _, child := range t.Children {
		culledChildren = append(culledChildren, child.Culled()...)
	}

	// Nodes to cull
	if strings.HasPrefix(t.Value, "#alternative") {
		return culledChildren
	}
	if t.Value == "#alternatives" {
		return culledChildren
	}
	if t.Value == "#item" {
		return culledChildren
	}
	if t.Value == "#items" {
		return culledChildren
	}

	return []Tree{
		{
			First:    t.First,
			Last:     t.Last,
			Value:    t.Value,
			Children: culledChildren,
		},
	}
}

func (t Tree) Flattened() Tree {
	var flat []Tree
	for _, child := range t.Children {
		if child.Value == t.Value {
			flat = append(flat, child.Flattened().Children...)
		} else {
			flat = append(flat, child.Flattened())
		}
	}

	cp := t
	cp.Children = flat
	return cp
}

func (t Tree) PrintIndent(indent int, s string, logger log.Logger) {
	logger.Log(fmt.Sprintf("%s(%d, %d) %s", strings.Repeat("  ", indent), t.First, t.Last, t.Value))

	if len(t.Children) == 0 {
		escaped := strings.ReplaceAll(s[t.First:t.Last], "\n", "\\n")
		logger.Log(fmt.Sprintf("= `%s`", escaped))
	}

	logger.Log("\n")

	for _, child := range t.Children {
		child.PrintIndent(indent+1, s, logger)
	}
}

func (t Tree) JsonSimplified() []Tree {
	if t.Value == "string" {
		t.Children = []Tree{}
		return []Tree{t}
	}
	if t.Value == "ws" {
		return []Tree{}
	}

	simplified := []Tree{}

	for _, child := range t.Children {
		simplified = append(simplified, child.JsonSimplified()...)
	}

	if t.Value == "element" {
		return simplified
	}
	if t.Value == "elements" {
		return simplified
	}
	if t.Value == "members" {
		return simplified
	}
	if t.Value == "value" {
		return simplified
	}

	t.Children = simplified
	return []Tree{t}
}

func (t Tree) McKeemanSimplified() []Tree {
	//if t.Value == "character" {
	//	return []Tree{}
	//}
	if t.Value == "hexcode" {
		t.Children = []Tree{}
		return []Tree{}
	}
	//if t.Value == "literal" {
	//	t.Children = []Tree{}
	//	return []Tree{t}
	//}
	if t.Value == "indentation" {
		return []Tree{}
	}
	if t.Value == "letter" {
		return []Tree{}
	}
	if t.Value == "newline" {
		return []Tree{}
	}
	if t.Value == "space" {
		return []Tree{}
	}

	simplified := []Tree{}

	for _, child := range t.Children {
		simplified = append(simplified, child.McKeemanSimplified()...)
	}

	if t.Value == "alternatives" {
		return simplified
	}
	//if t.Value == "item" {
	//	return simplified
	//}
	if t.Value == "items" {
		return simplified
	}
	if t.Value == "literal" {
		return simplified
	}
	if t.Value == "singleton" {
		return simplified
	}

	t.Children = simplified
	return []Tree{t}
}

func (t Tree) Go(src string) string {
	fields := ""
	children := ""
	for _, child := range t.Children {
		children += child.Go(src)
	}
	if children != "" {
		children = "Children: []*Tree{\n" + children + "},"
	} else {
		s := src[t.First:t.Last]
		s = strings.ReplaceAll(s, `\`, `\\`)
		s = strings.ReplaceAll(s, `"`, `\"`)
		//s = strings.ReplaceAll(s, "\n", `\n`)
		//s = strings.ReplaceAll(s, "\r", `\r`)
		//s = strings.ReplaceAll(s, "\t", `\t`)
		fields = "data: \"" + s + "\","
	}
	return fmt.Sprintf(
		`{Type: "%s", %s %s},`+"\n",
		t.Value,
		fields,
		children,
	)
}
