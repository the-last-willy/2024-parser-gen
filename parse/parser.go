package parse

import (
	"strconv"
	"strings"
)

type Parser struct {
	Grammar      *Tree
	GrammarStore string

	Rules map[string]*Tree
	Start *Tree
}

func NewParser(grammar *Tree, store string) *Parser {
	p := &Parser{}

	p.Grammar = grammar
	p.GrammarStore = store

	p.Start = grammar.FindFirst("rule")
	rules := grammar.FindAll("rule")

	p.Rules = make(map[string]*Tree)
	for _, rule := range rules {
		name := rule.FindFirst("name")
		p.Rules[name.GetData(p.GrammarStore)] = rule
	}

	return p
}

func (p *Parser) Parse(s string) *Tree {
	return p.ParseRule(p.Start, s, 0)
}

func (p *Parser) ParseRule(r *Tree, s string, cursor int) *Tree {
	if r.Type != "rule" {
		panic("not a rule")
	}
	name := r.FindFirst("name")
	alternatives := r.FindAll("alternative")
	// Must parse any alternative
	for _, alternative := range alternatives {
		items, ok := p.ParseAlternative(alternative, s, cursor)
		if !ok {
			continue
		}

		itemsWithoutPrimitives := []*Tree{}
		for _, item := range items {
			if item.Type != "xxx" {
				itemsWithoutPrimitives = append(itemsWithoutPrimitives, item)
			}
		}

		first := items[0].First
		last := items[len(items)-1].Last

		return &Tree{
			Type:     name.GetData(p.GrammarStore),
			First:    first,
			Last:     last,
			Children: itemsWithoutPrimitives,
		}
	}
	return nil
}

func (p *Parser) ParseAlternative(a *Tree, s string, cursor int) ([]*Tree, bool) {
	if a.Type != "alternative" {
		panic("not an alternative")
	}
	items := a.FindAll("item")
	parsedItems := []*Tree{}
	// Must parse all items
	for _, item := range items {
		parsed := p.ParseItem(item, s, cursor)
		if parsed == nil {
			return nil, false
		}
		cursor = parsed.Last
		parsedItems = append(parsedItems, parsed)
	}
	return parsedItems, true
}

func (p *Parser) ParseItem(i *Tree, s string, cursor int) *Tree {
	if i.Type != "item" {
		panic("not an item")
	}

	if nm := i.FindFirst("name"); nm != nil {
		return p.ParseRule(p.Rules[nm.GetData(p.GrammarStore)], s, cursor)
	}
	if cs := i.FindFirst("characters"); cs != nil {
		return p.ParseCharacters(cs, s, cursor)
	}
	if rn := i.FindFirst("range"); rn != nil {
		return p.ParseRangeItem(i, s, cursor)
	}

	if cp := i.FindFirst("codepoint"); cp != nil {
		return p.ParseCodePoint(cp, s, cursor)
	}

	panic("unknown item type")
}

// Parse primitives

func (p *Parser) ParseCharacters(c *Tree, s string, cursor int) *Tree {
	if c.Type != "characters" {
		panic("not a characters")
	}

	// Retrieves all characters
	data := ""
	characters := c.FindAll("character")
	if len(characters) == 0 {
		data = c.GetData(p.GrammarStore)
	} else {
		for _, char := range characters {
			data += char.GetData(p.GrammarStore)
		}
	}

	// Special case for empty characters
	if len(data) == 0 {
		return &Tree{
			Type:  "xxx",
			First: cursor,
			Last:  cursor,
		}
	}
	if strings.HasPrefix(s[cursor:], data) {
		return &Tree{
			Type:  "xxx",
			First: cursor,
			Last:  cursor + len(data),
		}
	}
	return nil
}

func (p *Parser) ParseCodePoint(c *Tree, s string, cursor int) *Tree {
	if c.Type != "codepoint" {
		panic("not codepoint")
	}
	if cursor >= len(s) {
		return nil
	}
	if s[cursor] == p.CodePointToUint8(c.GetData(p.GrammarStore)) {
		return &Tree{
			Type:  "xxx",
			First: cursor,
			Last:  cursor + 1,
		}
	}
	return nil
}

func (p *Parser) ParseRangeItem(item *Tree, s string, cursor int) *Tree {
	if item.Type != "item" {
		panic("ParseRangeItem: must receive the parent item")
	}
	if cursor >= len(s) {
		return nil
	}
	excludes := item.FindAll("exclude")
	for _, exclude := range excludes {
		if exclude.Type != "exclude" {
			panic("ParseRangeItem: not an exclude")
		}

		var t *Tree
		if cp := exclude.FindFirst("codepoint"); cp != nil {
			t = p.ParseCodePoint(cp, s, cursor)

		} else if rng := exclude.FindFirst("range"); rng != nil {
			panic("ParseRangeItem: range exclude not implemented")
		}
		// Exclusion should not be parsed
		if t != nil {
			return nil
		}
	}

	rang := item.FindFirst("range")
	if rang.Type != "range" {
		panic("ParseRangeItem: not a range")
	}

	cps := rang.FindAll("codepoint")
	lb := p.CodePointToUint8(cps[0].GetData(p.GrammarStore))
	ub := p.CodePointToUint8(cps[1].GetData(p.GrammarStore))
	if lb <= s[cursor] && s[cursor] <= ub {
		return &Tree{
			Type:  "xxx",
			First: cursor,
			Last:  cursor + 1,
		}
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
