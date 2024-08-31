package main

import (
	"fmt"
	"parsium/parse"
	"parsium/parse/lr"
	"parsium/tree"
	"slices"
)

func main() {
	parser := lr.Parser{
		ParseStack: nil,
		Cursor:     0,
	}
	_ = parser

	source := "A*2+1\x00"

	for _, r := range []rune(source) {
		parser.Process(r)
	}

	slices.Reverse(parser.ParseStack)

	tr := tree.NewSimpleTree2[parse.TreeData](
		&parse.TreeData{
			Type:  "root",
			First: 0,
			Last:  0,
		},
		parser.ParseStack...,
	)

	derec := parse.NewDerecursifier()
	_ = derec
	//tr = derec.Process(tr, tr).(tree.SimpleTree[parse.TreeData])

	formatter := parse.TreeFormatter{}
	fmt.Println(formatter.Format(tr, source))
}
