package json

import "parsium/parse"

func Simplify(tree parse.DataTree) parse.DataTree {
	return tree.WithRoot(simplify(tree)[0])
}

func simplify(tree parse.DataTree) []*parse.Tree {
	rootCopy := *tree.Root

	if rootCopy.Type == "ws" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{}
	}

	if rootCopy.Type == "number" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{&rootCopy}
	}
	if rootCopy.Type == "string" {
		rootCopy.Children = []*parse.Tree{}
		return []*parse.Tree{&rootCopy}
	}

	simplifiedChildren := []*parse.Tree{}
	for _, child := range tree.Root.Children {
		simplifiedChildren = append(simplifiedChildren, simplify(tree.WithRoot(child))...)
	}

	if rootCopy.Type == "element" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "members" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}
	if rootCopy.Type == "value" {
		rootCopy.Children = []*parse.Tree{}
		return simplifiedChildren
	}

	rootCopy.Children = simplifiedChildren
	return []*parse.Tree{&rootCopy}
}
