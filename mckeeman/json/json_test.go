package json

import (
	"parsium/assert"
	"testing"
)

func TestJson_NotAccept(t *testing.T) {
	g := JsonGrammar()
	assert.Equal(t, g.Accept(`"abcd""`), -1)
}
