package stack_test

import (
	"parsium/assert"
	"parsium/stack"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("At(0) returns first element of NewStack", func(t *testing.T) {
		s := stack.NewStack(7, 8, 9)
		assert.Equal(t, s.At(0), 7)
	})
	t.Run("At(0) returns second element before Pop", func(t *testing.T) {
		s := stack.NewStack(7, 8, 9)
		s = s.Pop()
		assert.Equal(t, s.At(0), 8)
	})
	t.Run("Len decrements after Pop", func(t *testing.T) {
		s := stack.NewStack(7, 8, 9)
		s = s.Pop()
		assert.Equal(t, s.Len(), 2)
	})
}
