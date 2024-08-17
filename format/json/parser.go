package json

import (
	"parsium/format/mckeeman"
	"parsium/parse"
)

type Parser struct {
	internal *mckeeman.Parser
}

func NewParser() *Parser {
	mkParser := mckeeman.NewParser()
	grammar := mkParser.Parse(mckeemanGrammar)

	return &Parser{
		internal: mckeeman.NewParserForGrammar(grammar),
	}
}

func (p *Parser) Parse(s string) parse.DataTree {
	return p.internal.Parse(s)
}
