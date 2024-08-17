package mck_parse

import (
	"fmt"
	"parsium/mckeeman"
)

type Parser struct {
	Grammar mckeeman.Grammar
	String  string

	Failed bool
}

func (p *Parser) Fail() {
	p.Failed = true
}

func (p *Parser) Status() string {
	if p.Failed {
		return "(failed)"
	}
	return ""
}

func Parse(g mckeeman.Grammar, s string) Tree {
	p := Parser{
		Grammar: g,
		String:  s,
	}
	t := p.ParseRule(
		g.FirstRule(),
		0,
	)
	if t.Last != len(s) {
		fmt.Println("FAILED TO PARSE AT", t.Last)
		fmt.Println(s[:t.Last], "XXX", s[t.Last:])
		p.Fail()
	}
	return t
}

// Parse functions

func (p *Parser) ParseAlternatives(as mckeeman.Alternatives, cursor int) Tree {
	ats := []Tree{}
	for _, a := range as.Alternatives {
		at := p.ParseAlternative(a, cursor)
		if !p.Failed {
			// Succeeded to parse one alternative
			return at
		}
		ats = append(ats, at)
		p.Failed = false
	}
	// Failed to parse any alternative
	p.Fail()
	return Tree{
		First:    cursor,
		Last:     cursor,
		Value:    "",
		Children: ats,
	}
}

func (p *Parser) ParseAlternative(a mckeeman.Alternative, cursor int) Tree {
	return p.ParseItems(a.Items, cursor)
}

func (p *Parser) ParseCharacters(c mckeeman.Characters, cursor int) Tree {
	cst := Tree{
		First:    cursor,
		Last:     cursor,
		Value:    "",
		Children: []Tree{},
	}
	for _, cp := range c.CodePoints {
		t := p.ParseCodePoint(cp, cst.Last)
		if p.Failed {
			break
		}
		cst.Last = t.Last
	}
	return cst
}

func (p *Parser) ParseCodePoint(c mckeeman.CodePoint, cursor int) Tree {
	if len(p.String) <= cursor {
		p.Fail()
		return Tree{
			First:    cursor,
			Last:     cursor,
			Value:    "",
			Children: []Tree{},
		}
	}
	if rune(p.String[cursor]) == c.Rune {
		return Tree{
			First:    cursor,
			Last:     cursor + 1,
			Value:    "",
			Children: []Tree{},
		}
	}
	p.Fail()
	return Tree{
		First:    cursor,
		Last:     cursor,
		Value:    "",
		Children: []Tree{},
	}
}

func (p *Parser) ParseItem(i mckeeman.Item, cursor int) Tree {
	switch i.(type) {
	case mckeeman.Name:
		return p.ParseRule(p.Grammar.RuleWithName(i.(mckeeman.Name)), cursor)
	case mckeeman.Literal:
		return p.ParseLiteral(i.(mckeeman.Literal), cursor)
	default:
		panic("Unknown item type")
	}
}

func (p *Parser) ParseItems(is mckeeman.Items, cursor int) Tree {
	tree := Tree{
		First:    cursor,
		Last:     cursor,
		Value:    "",
		Children: make([]Tree, 0),
	}
	for _, i := range is.Items {
		itree := p.ParseItem(i, tree.Last)
		if p.Failed {
			break
		}
		tree.Last = itree.Last
		if itree.Value != "" {
			tree.Children = append(tree.Children, itree)
		}
	}
	return tree
}

func (p *Parser) ParseLiteral(l mckeeman.Literal, cursor int) Tree {
	switch l.(type) {
	case mckeeman.Characters:
		return p.ParseCharacters(l.(mckeeman.Characters), cursor)
	case mckeeman.Range:
		return p.ParseRange(l.(mckeeman.Range), cursor)
	case mckeeman.Singleton:
		return p.ParseSingleton(l.(mckeeman.Singleton), cursor)
	default:
		panic("Unknown literal type")
	}
}

func (p *Parser) ParseRange(r mckeeman.Range, cursor int) Tree {
	if len(p.String) <= cursor {
		p.Fail()
		return Tree{
			First:    cursor,
			Last:     cursor,
			Value:    "",
			Children: []Tree{},
		}
	}
	if r.Accept(p.Grammar, p.String[cursor:]) != -1 {
		return Tree{
			First:    cursor,
			Last:     cursor + 1,
			Value:    "",
			Children: []Tree{},
		}
	}
	p.Fail()
	return Tree{
		First:    cursor,
		Last:     cursor,
		Value:    "",
		Children: []Tree{},
	}
}

func (p *Parser) ParseRule(r mckeeman.Rule, cursor int) Tree {
	ast := p.ParseAlternatives(r.Alternatives, cursor)
	ast.Value = r.Name().Value
	return ast
}

func (p *Parser) ParseSingleton(s mckeeman.Singleton, cursor int) Tree {
	return p.ParseCodePoint(s.CodePoint, cursor)
}
