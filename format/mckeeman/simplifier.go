package mckeeman

import (
	"parsium/parse"
	"parsium/tree"
	"slices"
)

func Simplify(tree parse.DataTree) parse.DataTree {
	return tree.WithRoot(simplify(tree)[0])
}

func simplify(tree parse.DataTree) []*parse.Tree {
	rootCopy := *tree.Root

	if rootCopy.Type == "indentation" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{}
	}
	if rootCopy.Type == "newline" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{}
	}
	if rootCopy.Type == "space" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{}
	}

	if rootCopy.Type == "name" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{&rootCopy}
	}

	simplifiedChildren := []*parse.Tree{}
	for _, child := range tree.Root.Children {
		simplifiedChildren = append(simplifiedChildren, simplify(tree.WithRoot(child))...)
	}

	if rootCopy.Type == "alternatives" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "element" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "elements" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "items" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "literal" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "rules" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "singleton" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}

	rootCopy.Children = simplifiedChildren
	return []*parse.Tree{&rootCopy}
}

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
