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
		return st.DataOf(*st.Root()).Type == ty
	}
}

func DataOf(st tree.SubTree[TreeData]) TreeData {
	return st.DataOf(*st.Root())
}

func TextOf(st tree.SubTree[TreeData], src string) string {
	d := DataOf(st)
	return src[d.First:d.Last]
}

func TypeOf(st tree.SubTree[TreeData]) string {
	return st.DataOf(*st.Root()).Type
}

func FindAllWithType(st tree.SubTree[TreeData], ty string) []tree.Node {
	return tree.FindAll(st, HasTypePred(ty))
}

func FindFirstWithType(st tree.SubTree[TreeData], ty string) *tree.Node {
	return tree.FindFirst(st, HasTypePred(ty))
}
