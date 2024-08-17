package mckeeman

import (
	"testing"
)

func TestCharacters_Accept_Same(t *testing.T) {
	c := NewCharacters("abcd")
	if c.Accept(Grammar{}, "abcd") != 4 {
		t.Error()
	}
}

func TestCharacters_Accept_Different(t *testing.T) {
	c := NewCharacters("abcd")
	if c.Accept(Grammar{}, "efgh") != -1 {
		t.Error()
	}
}

func TestCharacters_Accept_Empty(t *testing.T) {
	c := NewCharacters("")
	if c.Accept(Grammar{}, "abcd") != 0 {
		t.Error()
	}
}
