package lambda_calculus

func NormalEvaluation(e Expression) Expression {
	switch e := e.(type) {
	case Application:
		s := BetaReduce(e)
		if s == e {
			// Did not reduce, reduce function and argument.
			return ReuseApplication(e, BetaReduce(e.Function), BetaReduce(e.Argument))
		}
		// Did reduce, keep reducing.
		return BetaReduce(s)
	case Function:
		return e
	case Name:
		return e
	default:
		panic("EvaluateNormal: unhandled Expression type")
	}
}
