package mckeeman

import "fmt"

type Range struct {
	lb Singleton
	ub Singleton

	Exclude *Exclude
}

func NewRange(lb rune, ub rune) Literal {
	return Range{
		lb: NewSingleton(lb),
		ub: NewSingleton(ub),
	}
}

func NewRangeWithExclude(lb Singleton, ub Singleton, e Exclude) Range {
	return Range{lb: lb, ub: ub, Exclude: &e}
}

func NewRangeExclude(lb rune, ub rune, ex Exclude) Range {
	return Range{lb: NewSingleton(lb), ub: NewSingleton(ub), Exclude: &ex}
}

func (r Range) Accept(g Grammar, s string) int {
	if len(s) == 0 {
		return -1
	}
	if rune(s[0]) >= r.lb.CodePoint.Rune && rune(s[0]) <= r.ub.CodePoint.Rune {
		if r.Exclude != nil {
			return r.Exclude.Accept(g, s)
		}
		return 1
	}
	return -1
}

func (r Range) Print() {
	r.lb.Print()
	fmt.Print(" . ")
	r.ub.Print()
	if r.Exclude != nil {
		r.Exclude.Print()
	}
}
