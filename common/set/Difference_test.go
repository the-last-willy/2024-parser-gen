package set_test

import (
	"parsium/assert"
	"parsium/common/set"
	"testing"
)

func TestDifference(t *testing.T) {
	t.Run("works", func(t *testing.T) {
		a := set.NewSimpleSet(1, 2, 3)
		b := set.NewSimpleSet(2, 4)

		d := set.Difference[int](a, b)

		assert.True(t, d.Has(1))
		assert.True(t, !d.Has(2))
		assert.True(t, d.Has(3))
		assert.True(t, !d.Has(4))
	})
}
