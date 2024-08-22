package parse

import (
	"fmt"
	"parsium/tree"
	"strings"
)

type TreeData struct {
	Type string

	First int
	Last  int
}

func HasTypePred(ty string) func(tree.Tree[TreeData]) bool {
	return func(tr tree.Tree[TreeData]) bool {
		return tr.DataOf(tr.Root()).Type == ty
	}
}

func DataOf(t tree.Tree[TreeData]) TreeData {
	return t.DataOf(t.Root())
}

func TextOf(t tree.Tree[TreeData], src string) string {
	d := DataOf(t)
	return src[d.First:d.Last]
}

func TypeOf(t tree.Tree[TreeData]) string {
	return t.DataOf(t.Root()).Type
}

func FindAllWithType(t tree.Tree[TreeData], ty string) []tree.Node {
	return tree.FindAll(t, HasTypePred(ty))
}

func FindFirstWithType(t tree.Tree[TreeData], ty string) *tree.Node {
	return tree.FindFirst(t, HasTypePred(ty))
}

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
		panic("stored data are different")
	}
	return t.Data
}

func (t *Tree) PrintIndent(store string, indent int) {
	id := strings.Repeat("  ", indent)
	data := ""
	if len(t.Children) == 0 {
		data = store[t.First:t.Last]
		data = strings.ReplaceAll(data, "\n", `\n`)
		//data = strings.ReplaceAll(data, `"`, `\"`)
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

func (t *Tree) ToTreeNode(tr tree.Tree[TreeData]) tree.Node {
	children := []tree.Node{}
	for _, child := range t.Children {
		children = append(children, child.ToTreeNode(tr))
	}
	return tr.NewNode(
		TreeData{
			Type:  t.Type,
			First: t.First,
			Last:  t.Last,
		},
		children)
}
