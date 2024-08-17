package json

import (
	"parsium/assert"
	"testing"
)

func AssertValueAccept(t *testing.T, s string) {
	g := JsonGrammar()
	r := g.RuleWithNameFromString("value")
	assert.Equal(t, r.Accept(g, s), len(s))
}

func AssertValueNotAccept(t *testing.T, s string) {
	g := JsonGrammar()
	r := g.RuleWithNameFromString("value")
	assert.Equal(t, r.Accept(g, s), -1)
}

func TestValue_Array(t *testing.T) {
	AssertValueAccept(t, `[]`)
}

func TestValue_Empty(t *testing.T) {
	AssertValueNotAccept(t, ``)
}

func TestValue_False(t *testing.T) {
	AssertValueAccept(t, `false`)
}

func TestValue_Null(t *testing.T) {
	AssertValueAccept(t, `null`)
}

func TestValue_Number(t *testing.T) {
	AssertValueAccept(t, `0`)
}

func TestValue_Object(t *testing.T) {
	AssertValueAccept(t, `{}`)
}

func TestValue_String(t *testing.T) {
	AssertValueAccept(t, `""`)
}

func TestValue_True(t *testing.T) {
	AssertValueAccept(t, `true`)
}
