package tree_testhelpers

import (
	"parsium/assert"
	"parsium/tree"
	"testing"
)

func TestAsTree[TreeType tree.Tree[any]](t *testing.T, new func() TreeType) {
	t.Run("implements tree.Tree", func(t *testing.T) {
		assert.Implements[tree.Tree[any]](t, new())
	})
	t.Run("IsEmpty returns true for default", func(t *testing.T) {
		assert.True(t, new().IsEmpty())
	})
	t.Run("ChildrenOf returns node's children", func(test *testing.T) {
		t := new()
		child := t.NewNode(1, []tree.Node{})
		parent := t.NewNode(1, []tree.Node{child})
		assert.SliceEqual(test, t.ChildrenOf(parent), []tree.Node{child})
	})
	t.Run("DataOf returns node's data", func(test *testing.T) {
		t := new()
		n := t.NewNode(1, []tree.Node{})
		assert.Equal(test, t.DataOf(n), 1)
	})
	t.Run("Root returns WithRoot", func(test *testing.T) {
		var t tree.Tree[any] = new()
		n := t.NewNode(0, []tree.Node{})
		t = t.WithRoot(n)
		assert.Equal(test, t.Root(), n)
	})
}
