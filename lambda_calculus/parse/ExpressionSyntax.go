package lambda_calculus_parse

import (
	"parsium/common/set"
	"parsium/parse"
	"parsium/tree"
)

type ExpressionSyntax = tree.SimpleTree[parse.TreeData]

func ExpressionBoundVariables(e tree.Tree[parse.TreeData], source string) set.SimpleSet[string] {
	bvs := set.NewSimpleSet[string]()

	trav := tree.NewTraverser[parse.TreeData]()
	trav.Enter = func(s tree.Tree[parse.TreeData]) tree.TraverserCommand {
		if parse.TypeOf(s) == FunctionToken {
			bvs = bvs.Add(FunctionBoundVariable(s, source)).(set.SimpleSet[string])
		}

		return tree.TraverserContinue
	}

	trav.Process(e)

	return bvs
}

func ExpressionFreeVariables(e tree.Tree[parse.TreeData], source string) set.SimpleSet[string] {
	vars := set.NewSimpleSet[string]()

	trav := tree.NewTraverser[parse.TreeData]()
	trav.Enter = func(s tree.Tree[parse.TreeData]) tree.TraverserCommand {
		if parse.TypeOf(s) == NameToken {
			data := s.DataOf(*s.Root())
			v := source[data.First:data.Last]
			vars = vars.Add(v).(set.SimpleSet[string])
		}

		return tree.TraverserContinue
	}

	trav.Process(e)

	return set.Difference[string](vars, ExpressionBoundVariables(e, source))
}
