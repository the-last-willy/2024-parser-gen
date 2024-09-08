package lambda_calculus_parse

import (
	"parsium/parse"
	"parsium/tree"
)

func FunctionBoundVariable(e tree.Tree[parse.TreeData], src string) string {
	if parse.TypeOf(e) != FunctionToken {
		panic("FunctionBoundVariable: e is not a function")
	}

	bv := e.DataOf(e.ChildrenOf(*e.Root())[1])

	return src[bv.First:bv.Last]
}
