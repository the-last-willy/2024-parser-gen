package json

import (
	"parsium/assert"
	"testing"
)

func AssertArrayAccept(t *testing.T, s string) {
	g := JsonGrammar()
	r := g.RuleWithNameFromString("array")
	assert.Equal(t, r.Accept(g, s), len(s))
}

func TestArray_Empty(t *testing.T) {
	AssertArrayAccept(t, "[]")
}

func TestArray_Ws(t *testing.T) {
	AssertArrayAccept(t, "[ ]")
}

func TestArray_Nested(t *testing.T) {
	AssertArrayAccept(t, "[[]]")
}
