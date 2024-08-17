package json

import (
	"parsium/assert"
	"testing"
)

func AssertWsAccept(t *testing.T, s string) {
	g := JsonGrammar()
	r := g.RuleWithNameFromString("ws")
	assert.Equal(t, r.Accept(g, s), len(s))
}

func TestWs_Empty(t *testing.T) {
	AssertWsAccept(t, "")
}

func TestWs_Space(t *testing.T) {
	AssertWsAccept(t, " ")
}
