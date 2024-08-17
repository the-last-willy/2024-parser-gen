package parse

import (
	"fmt"
	"strings"
)

type Tree struct {
	Type string

	First int
	Last  int

	Data string

	Children []*Tree
}

func (t *Tree) Derecursified() *Tree {
	children := []*Tree{}
	for _, child := range t.Children {
		child := child.Derecursified()
		if child.Type == t.Type {
			children = append(children, child.Children...)
		} else {
			children = append(children, child)
		}
	}

	cp := *t
	cp.Children = children
	return &cp
}

func (t *Tree) FindAll(ty string) []*Tree {
	found := []*Tree{}
	t.TraverseAll(func(node *Tree) {
		if node.Type == ty {
			found = append(found, node)
		}
	})
	return found
}

func (t *Tree) FindFirst(ty string) *Tree {
	var found *Tree
	t.TraverseAll(func(node *Tree) {
		if found == nil && node.Type == ty {
			found = node
		}
	})
	return found
}

func (t *Tree) GetData(store string) string {
	if store == "" {
		return t.Data
	}
	if t.Data == "" {
		return store[t.First:t.Last]
	}
	if store[t.First:t.Last] != t.Data {
		panic("stored Data are different")
	}
	return t.Data
}

func (t *Tree) PrintIndent(store string, indent int) {
	id := strings.Repeat("  ", indent)
	data := ""
	if len(t.Children) == 0 {
		data = store[t.First:t.Last]
		data = strings.ReplaceAll(data, "\n", `\n`)
		//Data = strings.ReplaceAll(Data, `"`, `\"`)
		data = ` = '` + data + `'`
	}
	fmt.Printf("%s(%d, %d) %s%s\n", id, t.First, t.Last, t.Type, data)
	for _, child := range t.Children {
		child.PrintIndent(store, indent+1)
	}
}

func (t *Tree) TraverseAll(do func(*Tree)) {
	do(t)
	for _, child := range t.Children {
		child.TraverseAll(do)
	}
}

// McKeeman specific

func (t *Tree) McKeemanSimplified(src string) []*Tree {
	cp := *t

	//if t.Type == "codepoint" {
	//	cp.Data = src[cp.First:cp.Last]
	//	cp.Children = []*Tree{}
	//	return []*Tree{&cp}
	//}
	if t.Type == "indentation" {
		return []*Tree{}
	}
	//if t.Type == "literal" {
	//	cp.Data = src[cp.First:cp.Last]
	//	cp.Children = []*Tree{}
	//	return []*Tree{&cp}
	//}
	if t.Type == "name" {
		cp.Data = src[cp.First:cp.Last]
		cp.Children = []*Tree{}
		return []*Tree{&cp}
	}
	if t.Type == "newline" {
		return []*Tree{}
	}
	if t.Type == "space" {
		return []*Tree{}
	}

	simplified := []*Tree{}
	for _, child := range t.Children {
		simplified = append(simplified, child.McKeemanSimplified(src)...)
	}

	//if t.Type == "alternatives" {
	//	return simplified
	//}
	//if t.Type == "item" {
	//	return simplified
	//}
	//if t.Type == "items" {
	//	return simplified
	//}
	//if t.Type == "rules" {
	//	return simplified
	//}

	cp.Children = simplified
	return []*Tree{&cp}
}
