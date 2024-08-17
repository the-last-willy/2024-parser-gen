package mckeeman

type Grammar struct {
	Rules Rules
}

func NewGrammar(rules []Rule) Grammar {
	return Grammar{Rules: Rules{Rules: rules}}
}

// Accept returns the number of accepted characters.
// If s is not fully accepted, then it is not accepted.
func (g Grammar) Accept(s string) int {
	cnt := g.Rules.Accept(g, s)
	if cnt != len(s) {
		return -1
	}
	return cnt
}

func (g Grammar) FirstRule() Rule {
	return g.Rules.Rules[0]
}

func (g Grammar) Print() {
	g.Rules.Print()
}

func (g Grammar) RuleWithName(n Name) Rule {
	for _, r := range g.Rules.Rules {
		if r.name == n.Value {
			return r
		}
	}
	panic("rule not found")
}

func (g Grammar) RuleWithNameFromString(s string) Rule {
	return g.RuleWithName(NewName(s))
}
