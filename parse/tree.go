package parse

import (
	"parsium/tree"
)

type TreeData struct {
	Type string

	First int
	Last  int
}

func HasTypePred(ty string) func(tree.SubTree[TreeData]) bool {
	return func(st tree.SubTree[TreeData]) bool {
		return st.DataOf(st.Root()).Type == ty
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

func FindAllWithType(st tree.SubTree[TreeData], ty string) []tree.Node {
	return tree.FindAll(st, HasTypePred(ty))
}

func FindFirstWithType(st tree.SubTree[TreeData], ty string) *tree.Node {
	return tree.FindFirst(st, HasTypePred(ty))
}
