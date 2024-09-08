package tree_test

import (
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
