package lr

import (
	"parsium/parse"
	"parsium/tree"
	"unicode"
)

type data = parse.TreeData
type tr = tree.SimpleTree[data]

type Parser struct {
	ParseStack []tr
	LookAhead  tr
	Cursor     int
}

func (p *Parser) Process(r rune) {
	// Lexing

	if unicode.IsLetter(r) {
		if p.LookAhead.Root() != nil {
			p.ParseStack = append([]tr{p.LookAhead}, p.ParseStack...)
		}
		p.LookAhead = tree.NewSimpleTree2[data](
			&data{
				Type:  "id",
				First: p.Cursor,
				Last:  p.Cursor + 1,
			},
		)
	} else if unicode.IsDigit(r) {
		if p.LookAhead.Root() != nil {
			p.ParseStack = append([]tr{p.LookAhead}, p.ParseStack...)
		}
		p.LookAhead = tree.NewSimpleTree2[data](
			&data{
				Type:  "int",
				First: p.Cursor,
				Last:  p.Cursor + 1,
			},
		)
	} else {
		if p.LookAhead.Root() != nil {
			p.ParseStack = append([]tr{p.LookAhead}, p.ParseStack...)
		}
		p.LookAhead = tree.NewSimpleTree2[data](
			&data{
				Type:  string(r),
				First: p.Cursor,
				Last:  p.Cursor + 1,
			},
		)
	}

	p.Cursor += 1

	for {
		if p.Match("id", "id") {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "id",
					First: tree.RootData(p.ParseStack[1]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[1],
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[2:]...)
		} else if p.Match("int", "int") {
			n := tree.NewSimpleTree2[data](
				&data{
					Type:  "int",
					First: tree.RootData(p.ParseStack[1]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], n), p.ParseStack[2:]...)
		} else {
			break
		}
	}

	if len(p.ParseStack) == 0 {
		return
	}

	if parse.TypeOf(p.LookAhead) == parse.TypeOf(p.ParseStack[0]) {
		return
	}

	// Parsing

	for {
		if p.Match("Sums", "+", "Products") {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "Sums",
					First: tree.RootData(p.ParseStack[2]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[2],
				p.ParseStack[1],
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[3:]...)
		} else if p.Match("Products") && parse.TypeOf(p.LookAhead) != "*" {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "Sums",
					First: tree.RootData(p.ParseStack[0]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[1:]...)
		} else if p.Match("Products", "*", "Value") {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "Products",
					First: tree.RootData(p.ParseStack[2]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[2],
				p.ParseStack[1],
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[3:]...)
		} else if p.Match("Value") {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "Products",
					First: tree.RootData(p.ParseStack[0]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[1:]...)
		} else if p.Match("int") {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "Value",
					First: tree.RootData(p.ParseStack[0]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[1:]...)
		} else if p.Match("id") {
			s := tree.NewSimpleTree2[data](
				&data{
					Type:  "Value",
					First: tree.RootData(p.ParseStack[0]).First,
					Last:  tree.RootData(p.ParseStack[0]).Last,
				},
				p.ParseStack[0],
			)
			p.ParseStack = append(append(p.ParseStack[:0], s), p.ParseStack[1:]...)
		} else {
			break
		}
	}
}

func (p *Parser) Match(types ...string) bool {
	if len(types) > len(p.ParseStack) {
		return false
	}
	for i := range types {
		if types[len(types)-i-1] != tree.RootData(p.ParseStack[i]).Type {
			return false
		}
	}
	return true
}
