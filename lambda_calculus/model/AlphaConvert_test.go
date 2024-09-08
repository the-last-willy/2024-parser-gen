package lambda_calculus_test

import (
	"parsium/assert"
	lc "parsium/lambda_calculus/model"
	"testing"
)

func TestAlphaConvert(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		input := lc.Function{
			BoundVariable: lc.Name("a"),
			Expression:    lc.Name("a"),
		}
		output := lc.AlphaConvert(input, "a", "b").(lc.Function)
		expected := lc.Function{
			BoundVariable: lc.Name("b"),
			Expression:    lc.Name("b"),
		}
		assert.Equal(t, output, expected)
	})
}
