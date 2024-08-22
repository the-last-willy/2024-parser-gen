package mckeeman

import (
	"parsium/parse"
	"parsium/tree"
	"strconv"
	"strings"
)

type Parser struct {
	Grammar      tree.Tree[parse.TreeData]
	GrammarStore string

	// name to rule node
	Rules map[string]tree.Node
}

func (p *Parser) dataOf(n tree.Node) parse.TreeData {
	return parse.DataOf(p.Grammar.WithRoot(n))
}

func (p *Parser) textOf(n tree.Node) string {
	return parse.TextOf(p.Grammar.WithRoot(n), p.GrammarStore)
}

func (p *Parser) typeOf(n tree.Node) string {
	return parse.TypeOf(p.Grammar.WithRoot(n))
}

func (p *Parser) findAllWithType(n tree.Node, ty string) []tree.Node {
	return parse.FindAllWithType(p.Grammar.WithRoot(n), ty)
}

func (p *Parser) findFirstWithType(n tree.Node, ty string) *tree.Node {
	return parse.FindFirstWithType(p.Grammar.WithRoot(n), ty)
}

func (p *Parser) newLeaf(ty string, first, last int) tree.Node {
	return p.newNode(ty, first, last, []tree.Node{})
}

func (p *Parser) newNode(ty string, first, last int, children []tree.Node) tree.Node {
	return p.Grammar.NewNode(
		parse.TreeData{
			Type:  ty,
			First: first,
			Last:  last,
		},
		children,
	)
}

func NewParserForGrammar(grammar tree.Tree[parse.TreeData], store string) *Parser {
	p := &Parser{}

	p.Grammar = grammar
	p.GrammarStore = store

	rules := p.findAllWithType(grammar.Root(), "rule")

	p.Rules = map[string]tree.Node{}
	for _, rule := range rules {
		name := *p.findFirstWithType(rule, "name")
		nd := grammar.DataOf(name)
		p.Rules[p.GrammarStore[nd.First:nd.Last]] = rule
	}

	return p
}

func NewParser() *Parser {
	return NewParserForGrammar(grammarTree(), grammarSource)
}

func (p *Parser) Parse(s string) tree.Tree[parse.TreeData] {
	return p.Grammar.WithRoot(*p.ParseRule(*p.findFirstWithType(p.Grammar.Root(), "rule"), s, 0))
}

func (p *Parser) ParseRule(r tree.Node, s string, cursor int) *tree.Node {
	if p.typeOf(r) != "rule" {
		panic("not a rule")
	}
	name := *p.findFirstWithType(r, "name")
	alternatives := p.findAllWithType(r, "alternative")
	// Must parse any alternative
	for _, alternative := range alternatives {
		items, ok := p.ParseAlternative(alternative, s, cursor)
		if !ok {
			continue
		}

		itemsWithoutPrimitives := []tree.Node{}
		for _, item := range items {
			if p.typeOf(item) != "xxx" {
				itemsWithoutPrimitives = append(itemsWithoutPrimitives, item)
			}
		}

		first := p.dataOf(items[0]).First
		last := p.dataOf(items[len(items)-1]).Last

		nn := p.newNode(
			p.textOf(name),
			first,
			last,
			itemsWithoutPrimitives,
		)
		return &nn
	}
	return nil
}

func (p *Parser) ParseAlternative(a tree.Node, s string, cursor int) ([]tree.Node, bool) {
	if p.typeOf(a) != "alternative" {
		panic("not an alternative")
	}
	items := p.findAllWithType(a, "item")
	parsedItems := []tree.Node{}
	// Must parse all items
	for _, item := range items {
		parsed := p.ParseItem(item, s, cursor)
		if parsed == nil {
			return nil, false
		}
		cursor = p.dataOf(*parsed).Last
		parsedItems = append(parsedItems, *parsed)
	}
	return parsedItems, true
}

func (p *Parser) ParseItem(i tree.Node, s string, cursor int) *tree.Node {
	if p.typeOf(i) != "item" {
		panic("not an item")
	}

	if nm := p.findFirstWithType(i, "name"); nm != nil {
		r := p.Rules[p.textOf(*nm)]
		return p.ParseRule(r, s, cursor)
	}
	if cs := p.findFirstWithType(i, "characters"); cs != nil {
		return p.ParseCharacters(*cs, s, cursor)
	}
	if rn := p.findFirstWithType(i, "range"); rn != nil {
		return p.ParseRangeItem(i, s, cursor)
	}

	if cp := p.findFirstWithType(i, "codepoint"); cp != nil {
		return p.ParseCodePoint(*cp, s, cursor)
	}

	panic("unknown item type")
}

// Parse primitives

func (p *Parser) ParseCharacters(c tree.Node, s string, cursor int) *tree.Node {
	if p.typeOf(c) != "characters" {
		panic("not a characters")
	}

	// Retrieves all characters
	data := ""
	characters := p.findAllWithType(c, "character")
	if len(characters) == 0 {
		data = p.textOf(c)
	} else {
		for _, char := range characters {
			data += p.textOf(char)
		}
	}

	// Special case for empty characters
	if len(data) == 0 {
		leaf := p.newLeaf("xxx", cursor, cursor)
		return &leaf
	}
	if strings.HasPrefix(s[cursor:], data) {
		leaf := p.newLeaf("xxx", cursor, cursor+len(data))
		return &leaf
	}
	return nil
}

func (p *Parser) ParseCodePoint(c tree.Node, s string, cursor int) *tree.Node {
	if p.typeOf(c) != "codepoint" {
		panic("not codepoint")
	}
	if cursor >= len(s) {
		return nil
	}
	if s[cursor] == p.CodePointToUint8(p.textOf(c)) {
		leaf := p.newLeaf("xxx", cursor, cursor+1)
		return &leaf
	}
	return nil
}

func (p *Parser) ParseRangeItem(item tree.Node, s string, cursor int) *tree.Node {
	if p.typeOf(item) != "item" {
		panic("ParseRangeItem: must receive the parent item")
	}
	if cursor >= len(s) {
		return nil
	}
	excludes := p.findAllWithType(item, "exclude")
	for _, exclude := range excludes {
		if p.typeOf(exclude) != "exclude" {
			panic("ParseRangeItem: not an exclude")
		}

		var t *tree.Node
		if cp := p.findFirstWithType(exclude, "codepoint"); cp != nil {
			t = p.ParseCodePoint(*cp, s, cursor)

		} else if rng := p.findFirstWithType(exclude, "range"); rng != nil {
			panic("ParseRangeItem: range exclude not implemented")
		}
		// Exclusion should not be parsed
		if t != nil {
			return nil
		}
	}

	rang := *p.findFirstWithType(item, "range")
	if p.typeOf(rang) != "range" {
		panic("ParseRangeItem: not a range")
	}

	cps := p.findAllWithType(rang, "codepoint")
	lb := p.CodePointToUint8(p.textOf(cps[0]))
	ub := p.CodePointToUint8(p.textOf(cps[1]))
	if lb <= s[cursor] && s[cursor] <= ub {
		leaf := p.newLeaf("xxx", cursor, cursor+1)
		return &leaf
	}
	return nil
}

func (p *Parser) CodePointToUint8(s string) uint8 {
	// TODO Hex weirdness, encoding...
	if len(s) == 1 {
		return s[0]
	}

	i, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		panic(err)
	}
	if i >= 127 {
		return 127
	}
	return uint8(i)
}
