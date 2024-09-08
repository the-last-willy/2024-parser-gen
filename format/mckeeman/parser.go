package mckeeman

import (
	"parsium/parse"
	"parsium/tree"
	"strconv"
	"strings"
)

type parseResult struct {
	last  int
	nodes []tree.Builder[parse.TreeData]
}

func (r parseResult) Single() tree.Builder[parse.TreeData] {
	if len(r.nodes) == 1 {
		return r.nodes[0]
	}
	panic("parseResult.Single")
}

type Parser struct {
	Grammar      tree.Tree[parse.TreeData]
	GrammarStore string

	Builder tree.Builder[parse.TreeData]

	// name to rule node
	Rules map[string]tree.Node

	alternativeToItems map[tree.Node][]tree.Node
	ruleToAlternatives map[tree.Node][]tree.Node
	ruleToName         map[tree.Node]string
}

func (p *Parser) textOf(n tree.Node) string {
	return parse.TextOf(tree.NewSubTree(p.Grammar, &n), p.GrammarStore)
}

func (p *Parser) typeOf(n tree.Node) string {
	return parse.TypeOf(tree.NewSubTree(p.Grammar, &n))
}

func (p *Parser) findAllWithType(n tree.Node, ty string) []tree.Node {
	return parse.FindAllWithType(tree.NewSubTree(p.Grammar, &n), ty)
}

func (p *Parser) findFirstWithType(n tree.Node, ty string) *tree.Node {
	return parse.FindFirstWithType(tree.NewSubTree(p.Grammar, &n), ty)
}

func NewParserForGrammar(grammar tree.Tree[parse.TreeData], store string) *Parser {
	p := &Parser{}

	p.Grammar = grammar
	p.GrammarStore = store

	rules := p.findAllWithType(*grammar.Root(), RuleType)

	p.Rules = map[string]tree.Node{}
	for _, rule := range rules {
		name := *p.findFirstWithType(rule, NameType)
		nd := grammar.DataOf(name)
		p.Rules[p.GrammarStore[nd.First:nd.Last]] = rule
	}

	p.alternativeToItems = map[tree.Node][]tree.Node{}
	for _, a := range p.findAllWithType(*p.Grammar.Root(), AlternativeType) {
		p.alternativeToItems[a] = p.findAllWithType(a, ItemType)
	}

	p.ruleToAlternatives = map[tree.Node][]tree.Node{}
	for _, r := range p.Rules {
		alternatives := p.findAllWithType(r, AlternativeType)
		p.ruleToAlternatives[r] = alternatives
	}

	p.ruleToName = map[tree.Node]string{}
	for _, r := range p.Rules {
		name := p.findFirstWithType(r, NameType)
		p.ruleToName[r] = p.textOf(*name)
	}

	return p
}

func NewParser() *Parser {
	return NewParserForGrammar(GrammarTree(), GrammarSource)
}

func (p *Parser) Parse(s string, b tree.Builder[parse.TreeData]) tree.Builder[parse.TreeData] {
	p.Builder = b
	start := p.findFirstWithType(*p.Grammar.Root(), RuleType)
	parsed := p.ParseRule(*start, s, 0)
	return parsed.Single()
}

func (p *Parser) ParseRule(r tree.Node, s string, cursor int) *parseResult {
	if p.typeOf(r) != RuleType {
		panic("not a rule")
	}

	// Must parse any alternative
	for _, alternative := range p.ruleToAlternatives[r] {
		alt := p.ParseAlternative(alternative, s, cursor)

		// Failed to parse alternative try next one
		if alt == nil {
			continue
		}

		// Succeeded to parse one alternative
		// TODO Incomplete

		return &parseResult{
			last: alt.last,
			nodes: []tree.Builder[parse.TreeData]{
				p.Builder.Tree(
					parse.TreeData{
						Type:  p.ruleToName[r],
						First: cursor,
						Last:  alt.last,
					},
					alt.nodes,
				),
			},
		}
	}

	// Failed to parse all alternatives
	return nil
}

func (p *Parser) ParseAlternative(a tree.Node, s string, cursor int) *parseResult {
	if p.typeOf(a) != AlternativeType {
		panic("not an alternative")
	}

	res := parseResult{
		last:  cursor,
		nodes: nil,
	}

	// Must parse all items
	for _, item := range p.alternativeToItems[a] {
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

		if cp := p.findFirstWithType(exclude, CodePointType); cp != nil {
			_, ok := p.ParseCodePoint(*cp, s, cursor)
			if ok {
				// Exclusion should not be parsed
				return 0, false
			}

		} else if rng := p.findFirstWithType(exclude, "range"); rng != nil {
			panic("ParseRangeItem: range exclude not implemented")
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
