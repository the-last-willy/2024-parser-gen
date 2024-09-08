package lambda_calculus

import (
	"parsium/assert"
	"testing"
)

func TestBetaReduce(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		app := Application{
			Function: Function{
				BoundVariable: Name("a"),
				Expression:    Name("a"),
			},
			Argument: Name("b"),
		}
		res := BetaReduce(app)
		var expected Expression = Name("b")
		assert.Equal(t, res, expected)
	})
}
