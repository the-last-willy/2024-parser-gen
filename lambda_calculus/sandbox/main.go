package main

import (
	"fmt"
	"parsium/common"
	"parsium/common/set"
	parse2 "parsium/lambda_calculus/parse"
	"parsium/parse"
	"parsium/tree"
)

func main() {
	src := `\  a.\b.c`

	lexer := parse2.NewLexer()
	parser := parse2.NewParser()

	builders := []tree.SimpleTree[parse.TreeData]{}

	lexer.Output = func(token parse.TreeData) {
		if token.Type != "space" {
			builders = append(builders, tree.NewSimpleTree2(&token))
			parser.Process(token.Type)
		}
	}
	parser.Output = func(token parse.TreeData) {
		builders = common.SliceReduceRange(builders, token.First, token.Last, func(ran []tree.SimpleTree[parse.TreeData]) []tree.SimpleTree[parse.TreeData] {
			subs := []tree.Builder[parse.TreeData]{}
			for _, b := range ran {
				subs = append(subs, b.ConvertTo(b))
			}
			data := token
			data.First = ran[0].DataOf(*ran[0].Root()).First
			last := common.SliceLast(ran)
			data.Last = last.DataOf(*last.Root()).Last
			return []tree.SimpleTree[parse.TreeData]{ran[0].Tree(data, subs).(tree.SimpleTree[parse.TreeData])}
		})
	}

	for _, c := range []byte(src) {
		lexer.Process(c)
	}
	lexer.Flush()
	parser.Flush()

	finalbuilder := tree.NewSimpleTree2(
		&parse.TreeData{
			Type:  "root",
			First: 0,
			Last:  0,
		},
		builders...,
	)

	tr := finalbuilder.Build()

	formatter := parse.TreeFormatter{}
	fmt.Println(formatter.Format(tr, src))

	fmt.Println(set.ToSlice[string](parse2.ExpressionBoundVariables(tr, src)))
	fmt.Println(set.ToSlice[string](parse2.ExpressionFreeVariables(tr, src)))
}
