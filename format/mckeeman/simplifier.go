package mckeeman

import (
	"parsium/parse"
	"parsium/tree"
	"slices"
)

type Simplifier struct {
	derecursifier *parse.Derecursifier
	pruner        *tree.Pruner[parse.TreeData]
}

func NewSimplifier() *Simplifier {
	return &Simplifier{
		derecursifier: parse.NewDerecursifier(),
		pruner: &tree.Pruner[parse.TreeData]{
			Remove: func(t tree.Tree[parse.TreeData], n tree.Node) bool {
				return slices.Contains(
					[]string{
						HexType,
						IndentationType,
						NewlineType,
						SpaceType,
					},
					t.DataOf(n).Type,
				)
			},
			RemoveChildren: func(t tree.Tree[parse.TreeData], n tree.Node) bool {
				return slices.Contains(
					[]string{
						NameType,
					},
					t.DataOf(n).Type,
				)
			},
			RemoveKeepChildren: func(t tree.Tree[parse.TreeData], n tree.Node) bool {
				return slices.Contains(
					[]string{
						AlternativesType,
						ItemsType,
						LiteralType,
						RulesType,
						SingletonType,
					},
					t.DataOf(n).Type,
				)
			},
		},
	}
}

func (s *Simplifier) Process(t tree.Tree[parse.TreeData]) tree.Tree[parse.TreeData] {
	t = s.derecursifier.Process(t)
	t = s.pruner.Apply(t)
	return t
}
