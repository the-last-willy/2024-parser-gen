package lambda_calculus_test

import (
	model "parsium/lambda_calculus/model"
	"testing"
)

func TestAlphaConvert(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		input := model.Function{
			BoundVariable: "",
			Expression:    nil,
		}
	})
}
