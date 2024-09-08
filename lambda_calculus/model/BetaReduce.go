package lambda_calculus

func BetaReduce(e Expression) Expression {
	if _, isApp := e.(Application); !isApp {
		return e
	}
	a := e.(Application)
	switch f := e.(Application).Function.(type) {
	case Function:
		return Substitute(f.Expression, f.BoundVariable, a.Argument)
	case Application:
		return e
	case Name:
		return e
	default:
		panic("BetaReduce: unhandled Expression type")
	}
}
