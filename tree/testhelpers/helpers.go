package tree_testhelpers

import (
	"parsium/tree"
	"testing"
)

func TestAsTree[TreeType tree.Tree[any]](t *testing.T, new func() TreeType) {
	//t.Run("implements tree.New", func(t *testing.T) {
	//	assert.Implements[tree.Tree[any]](t, new())
	//})
	//t.Run("ChildrenOf returns node's children", func(test *testing.T) {
	//	t := new()
	//	child := t.NewNode(1, []tree.Node{})
	//	parent := t.NewNode(1, []tree.Node{child})
	//	_ = parent
	//	//assert.SliceEqual(test, t.ChildrenOf(parent), []tree.New{child})
	//})
	//t.Run("DataOf returns node's data", func(test *testing.T) {
	//	t := new()
	//	n := t.NewNode(1, []tree.Node{})
	//	assert.Equal(test, t.DataOf(n), 1)
	//})
}
