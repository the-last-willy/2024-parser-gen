package ebnf

import (
	"parsium/parse"
	"parsium/tree"
	"slices"
	"strings"
)

func Simplify(t tree.Tree[parse.Tree]) tree.Tree[parse.Tree] {
	p := tree.Pruner[parse.Tree]{
		Remove: func(t tree.Tree[parse.Tree], n tree.Node) bool {
			return slices.Contains(
				[]string{
					"S",
				},
				t.DataOf(n).Type,
			)
		},
		RemoveKeepChildren: func(t tree.Tree[parse.Tree], n tree.Node) bool {
			if strings.HasPrefix(t.DataOf(n).Type, "x") {
				return true
			}
			return slices.Contains(
				[]string{
					"lhs",
					"rhs",
				},
				t.DataOf(n).Type,
			)
		},
		RemoveChildren: func(t tree.Tree[parse.Tree], n tree.Node) bool {
			return slices.Contains(
				[]string{
					"identifier",
				},
				t.DataOf(n).Type,
			)
		},
	}
	return p.Apply(t)
}
