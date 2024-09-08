package set

import (
	"parsium/assert"
	"testing"
)

func TestSimpleSet(t *testing.T) {
	t.Run("Has is false for NewSimpleSet", func(t *testing.T) {
		s := NewSimpleSet[int]()
		assert.False(t, s.Has(9))
	})
	t.Run("Has is true for NewSimpleSet with elements", func(t *testing.T) {
		s := NewSimpleSet[int](9)
		assert.True(t, s.Has(9))
	})
	t.Run("Has is true after Add", func(t *testing.T) {
		s := NewSimpleSet[int]()
		s = s.Add(9).(SimpleSet[int])
		assert.True(t, s.Has(9))
	})
	t.Run("Has is false after Remove", func(t *testing.T) {
		s := NewSimpleSet[int](9)
		s = s.Remove(9).(SimpleSet[int])
		assert.False(t, s.Has(9))
	})
}
