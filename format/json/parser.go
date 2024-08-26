package json

import (
	"parsium/format/mckeeman"
	"parsium/parse"
	"parsium/tree"
)

type Parser struct {
	internal *mckeeman.Parser
}

func NewParser() *Parser {
	mkParser := mckeeman.NewParser()
	grammar := mkParser.Parse(mckeemanGrammar, tree.NewSimpleTree[parse.TreeData]()).(tree.SimpleTree[parse.TreeData])

	return &Parser{
		internal: mckeeman.NewParserForGrammar(grammar, mckeemanGrammar),
	}
}

func (p *Parser) Parse(s string) tree.Tree[parse.TreeData] {
	return p.internal.Parse(s, tree.NewSimpleTree[parse.TreeData]()).(tree.SimpleTree[parse.TreeData])
}
