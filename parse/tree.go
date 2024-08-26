package parse

import (
	"parsium/tree"
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

func TypeOf(t tree.Tree[TreeData], n tree.Node) string {
	return t.DataOf(n).Type
}

func FindAllWithType(t tree.Tree[TreeData], ty string) []tree.Node {
	return tree.FindAll(t, HasTypePred(ty))
}

func FindFirstWithType(t tree.Tree[TreeData], ty string) *tree.Node {
	return tree.FindFirst(t, HasTypePred(ty))
}
