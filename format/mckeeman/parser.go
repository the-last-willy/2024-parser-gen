package mckeeman

import (
	"parsium/parse"
	"parsium/tree"
	"strconv"
	"strings"
)

type parseResult struct {
	last  int
	nodes []tree.Node
}

func (r parseResult) Single() tree.Node {
	if len(r.nodes) == 1 {
		return r.nodes[0]
	}
	panic("parseResult.Single")
}

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

	rules := p.findAllWithType(grammar.Root(), RuleType)

	p.Rules = map[string]tree.Node{}
	for _, rule := range rules {
		name := *p.findFirstWithType(rule, NameType)
		nd := grammar.DataOf(name)
		p.Rules[p.GrammarStore[nd.First:nd.Last]] = rule
	}

	return p
}

func NewParser() *Parser {
	return NewParserForGrammar(grammarTree(), grammarSource)
}

func (p *Parser) Parse(s string) tree.Tree[parse.TreeData] {
	start := p.findFirstWithType(p.Grammar.Root(), RuleType)
	parsed := p.ParseRule(*start, s, 0)
	return p.Grammar.WithRoot(parsed.nodes[0])
}

func (p *Parser) ParseRule(r tree.Node, s string, cursor int) *parseResult {
	if p.typeOf(r) != RuleType {
		panic("not a rule")
	}
	name := *p.findFirstWithType(r, NameType)
	alternatives := p.findAllWithType(r, AlternativeType)
	// Must parse any alternative
	for _, alternative := range alternatives {
		alt := p.ParseAlternative(alternative, s, cursor)
		if alt == nil {
			continue
		}

		nn := p.newNode(
			p.textOf(name),
			cursor,
			alt.last,
			alt.nodes,
		)

		return &parseResult{
			last:  alt.last,
			nodes: []tree.Node{nn},
		}
	}
	return nil
}

func (p *Parser) ParseAlternative(a tree.Node, s string, cursor int) *parseResult {
	if p.typeOf(a) != AlternativeType {
		panic("not an alternative")
	}
	items := p.findAllWithType(a, ItemType)

	res := parseResult{
		last:  cursor,
		nodes: nil,
	}

	// Must parse all items
	for _, item := range items {
		res2 := p.ParseItem(item, s, res.last)
		if res2 == nil {
			return nil
		}
		res.last = res2.last
		res.nodes = append(res.nodes, res2.nodes...)
	}

	return &res
}

func (p *Parser) ParseItem(i tree.Node, s string, cursor int) *parseResult {
	if p.typeOf(i) != ItemType {
		panic("not an item")
	}

	if nm := p.findFirstWithType(i, NameType); nm != nil {
		r := p.Rules[p.textOf(*nm)]
		return p.ParseRule(r, s, cursor)
	}

	var count int
	var ok bool

	if cs := p.findFirstWithType(i, CharactersType); cs != nil {
		count, ok = p.ParseCharacters(*cs, s, cursor)
	} else if rn := p.findFirstWithType(i, RangeType); rn != nil {
		count, ok = p.ParseRangeItem(i, s, cursor)
	} else if cp := p.findFirstWithType(i, CodePointType); cp != nil {
		count, ok = p.ParseCodePoint(*cp, s, cursor)
	} else {
		panic("unknown item type")
	}

	if ok {
		return &parseResult{
			last:  cursor + count,
			nodes: nil,
		}
	} else {
		return nil
	}
}

// Parse primitives

func (p *Parser) ParseCharacters(c tree.Node, s string, cursor int) (count int, ok bool) {
	if p.typeOf(c) != CharactersType {
		panic("not a characters")
	}

	// Retrieves all characters
	data := p.textOf(c)

	// Special case for empty characters
	if len(data) == 0 {
		return 0, true
	}
	if strings.HasPrefix(s[cursor:], data) {
		return len(data), true
	}
	return 0, false
}

func (p *Parser) ParseCodePoint(c tree.Node, s string, cursor int) (count int, ok bool) {
	if p.typeOf(c) != CodePointType {
		panic("not codepoint")
	}
	if cursor >= len(s) {
		return 0, false
	}
	if s[cursor] == p.CodePointToUint8(p.textOf(c)) {
		return 1, true
	}
	return 0, false
}

func (p *Parser) ParseRangeItem(item tree.Node, s string, cursor int) (count int, ok bool) {
	if p.typeOf(item) != ItemType {
		panic("ParseRangeItem: must receive the parent item")
	}
	if cursor >= len(s) {
		return 0, false
	}
	excludes := p.findAllWithType(item, ExcludeType)
	for _, exclude := range excludes {
		if p.typeOf(exclude) != ExcludeType {
			panic("ParseRangeItem: not an exclude")
		}

		var t *tree.Node
		if cp := p.findFirstWithType(exclude, CodePointType); cp != nil {
			count, ok := p.ParseCodePoint(*cp, s, cursor)
			if ok {
				l := p.newLeaf("xxx", cursor, cursor+count)
				t = &l
			}

		} else if rng := p.findFirstWithType(exclude, "range"); rng != nil {
			panic("ParseRangeItem: range exclude not implemented")
		}
		// Exclusion should not be parsed
		if t != nil {
			return 0, false
		}
	}

	rang := *p.findFirstWithType(item, "range")
	if p.typeOf(rang) != RangeType {
		panic("ParseRangeItem: not a range")
	}

	cps := p.findAllWithType(rang, CodePointType)
	lb := p.CodePointToUint8(p.textOf(cps[0]))
	ub := p.CodePointToUint8(p.textOf(cps[1]))
	if lb <= s[cursor] && s[cursor] <= ub {
		return 1, true
	}
	return 0, false
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
