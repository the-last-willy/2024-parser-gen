package main

import (
	"fmt"
	"parsium/parse"
	"parsium/parse/lr"
	"parsium/tree"
	"slices"
)

func main1() {
	//func main() {
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

func main2() {
	//func main() {
	rules := [][]string{
		{"Goal", "Sums", "eof"},
		{"Sums", "Sums", "+", "Products"},
		{"Sums", "Products"},
		{"Products", "Products", "*", "Value"},
		{"Products", "Value"},
		{"Value", "int"},
		{"Value", "id"},
	}

	graph := parse.NewGraph()
	graph.AddProduction(graph.Start, "start")

	for _, rule := range rules {
		cur := graph.Start

		for i := len(rule) - 1; i >= 1; i-- {
			cur = graph.AddTransition(cur, rule[i])
		}

		graph.AddProduction(cur, rule[0])
	}

	for v := range graph.Topology.Vs() {
		name, has := graph.Productions[v]
		if !has {
			name = fmt.Sprint(v)
		}
		fmt.Println(name)
	}
	fmt.Println()
	for from, transition_to := range graph.Transitions {
		for transition, to := range transition_to {
			v1name := graph.Productions[from]
			v1name += fmt.Sprint(from)
			v2name := graph.Productions[to]
			v2name += fmt.Sprint(to)
			fmt.Println(v1name, v2name, transition)
		}
	}
}

func main() {
	parser := parse.Lexer{}

}
