package tree_test

import (
	"fmt"
	"parsium/assert"
	"parsium/tree"
	tree_testhelpers "parsium/tree/testhelpers"
	"testing"
)

func TestTree(t *testing.T) {
	new := func() tree.SimpleTree[any] {
		return tree.SimpleTree[any]{}
	}
	tree_testhelpers.TestAsTree[tree.SimpleTree[any]](t, new)
}

func TestSimpleTree(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tr := tree.NewSimpleTree[int]().AsTree()
		tr = tr.WithRoot(tr.NewNode(3, []tree.Node{}))
		assert.Equal(t, tr.Root(), tr.Root())
	})
	t.Run("test2", func(t *testing.T) {
		b := tree.NewSimpleBuilder[int]().Tree(1, []tree.Builder[int]{
			tree.NewSimpleBuilder[int]().Tree(2, []tree.Builder[int]{}),
			tree.NewSimpleBuilder[int]().Tree(3, []tree.Builder[int]{}),
		}).(tree.SimpleBuilder[int])
		tr := b.Build()
		fmt.Printf("-%p\n", tr.ChildrenOf(tr.Root())[0].Impl)
		fmt.Printf("-%p\n", tr.ChildrenOf(tr.Root())[0].Impl)
	})
}
