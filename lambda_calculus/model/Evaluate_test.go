package lambda_calculus

import (
	"parsium/assert"
	"testing"
)

func TestNormalEvaluation(t *testing.T) {
	t.Run("", func(t *testing.T) {
		expr := NewApplication(
			NewFunction(
				NewName("a"),
				NewName("a"),
			),
			NewName("b"),
		)
		res := NormalEvaluation(expr)
		expected := NewName("b")
		assert.Equal[Expression](t, expected, res)
	})
}
