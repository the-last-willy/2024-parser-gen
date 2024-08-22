package json

import (
	"parsium/format/mckeeman"
	"parsium/parse"
	"parsium/tree"
)

type Parser struct {
	internal *mckeeman.BetterParser
}

func NewParser() *Parser {
	mkParser := mckeeman.NewBetterParser()
	grammar := mkParser.Parse(mckeemanGrammar)

	return &Parser{
		internal: mckeeman.NewBetterParserForGrammar(grammar, mckeemanGrammar),
	}
}

func (p *Parser) Parse(s string) tree.Tree[parse.TreeData] {
	return p.internal.Parse(s)
}
