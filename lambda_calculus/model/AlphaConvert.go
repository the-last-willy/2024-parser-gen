package lambda_calculus

func AlphaConvert(e Expression, before, after string) Expression {
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

func alphaConvertApplication(a Application, before, after string) Application {
	fun := AlphaConvert(a.Function, before, after)
	arg := AlphaConvert(a.Argument, before, after)

	if a.Function == fun && a.Argument == arg {
		return a
	}

	return Application{
		Function: fun,
		Argument: arg,
	}
}

func alphaConvertFunction(f Function, before, after string) Function {
	bv := f.BoundVariable

	expr := AlphaConvert(f.Expression, before, after)

	if f.BoundVariable == bv && f.Expression == expr {
		return f
	}

	return Function{
		BoundVariable: bv,
		Expression:    expr,
	}
}

func alphaConvertName(n Name, before, after string) Name {
	if string(n) == after {
		panic("AlphaConvert: name collision")
	}
	if string(n) == before {
		return Name(after)
	}
	return n
}
