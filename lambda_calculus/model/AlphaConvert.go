package lambda_calculus

func AlphaConvert(e Expression, before, after Name) Expression {
	switch e.(type) {
	case Application:
		return alphaConvertApplication(e.(Application), before, after)
	case Function:
		return alphaConvertFunction(e.(Function), before, after)
	case Name:
		return alphaConvertName(e.(Name), before, after)
	default:
		panic("AlphaConvert: unhandled Expression type")
	}
}

func alphaConvertApplication(a Application, before, after Name) Application {
	fun := AlphaConvert(a.Function, before, after)
	arg := AlphaConvert(a.Argument, before, after)
	return ReuseApplication(a, fun, arg)
}

func alphaConvertFunction(f Function, before, after Name) Function {
	bv := alphaConvertName(f.BoundVariable, before, after)
	expr := AlphaConvert(f.Expression, before, after)
	return ReuseFunction(f, bv, expr)
}

func alphaConvertName(n Name, before, after Name) Name {
	if n == after {
		panic("AlphaConvert: name collision")
	}
	if n == before {
		return Name(after)
	}
	return n
}
