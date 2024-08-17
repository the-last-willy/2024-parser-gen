package mckeeman

import "testing"

func TestCodePoint_Accept_Same(t *testing.T) {
	c := CodePoint{Rune: 'a'}
	if c.Accept("a") != 1 {
		t.Error()
	}
}

func TestCodePoint_Accept_Different(t *testing.T) {
	c := CodePoint{Rune: 'a'}
	if c.Accept("b") != -1 {
		t.Error()
	}
}

func TestCodePoint_Accept_Empty(t *testing.T) {
	c := CodePoint{Rune: 'a'}
	if c.Accept("") != -1 {
		t.Error()
	}
}
