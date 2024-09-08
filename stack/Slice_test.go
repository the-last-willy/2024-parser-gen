package stack_test

import (
	"parsium/assert"
	"parsium/stack"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("Slice(int, int) returns middle of slice", func(t *testing.T) {
		st := stack.NewStack(7, 8, 9)
		sl := stack.Slice(st, 1, 2)
		assert.SliceEqual(t, sl, []int{8})
	})
	t.Run("Slice(int, nil) returns end of slice", func(t *testing.T) {
		st := stack.NewStack(7, 8, 9)
		sl := stack.Slice(st, 1, nil)
		assert.SliceEqual(t, sl, []int{8, 9})
	})
	t.Run("Slice(nil, int) returns beginning of slice", func(t *testing.T) {
		st := stack.NewStack(7, 8, 9)
		sl := stack.Slice(st, nil, 2)
		assert.SliceEqual(t, sl, []int{7, 8})
	})
}
