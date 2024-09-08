package lambda_calculus

func Substitute(e Expression, before Name, after Expression) Expression {
	switch e.(type) {
	case Application:
		return substituteApplication(e.(Application), before, after)
	case Function:
		return substituteFunction(e.(Function), before, after)
	case Name:
		return substituteName(e.(Name), before, after)
	default:
		panic("Substitute: unhandled Expression type")
	}
}

func substituteApplication(a Application, before Name, after Expression) Application {
	fun := Substitute(a.Function, before, after)
	arg := Substitute(a.Argument, before, after)

	if a.Function == fun && a.Argument == arg {
		return a
	}

	return Application{
		Function: fun,
		Argument: arg,
	}
}

func substituteFunction(f Function, before Name, after Expression) Function {
	// Shadowing
	if f.BoundVariable == before {
		return f
	}

	expr := Substitute(f.Expression, before, after)

	if expr == f.Expression {
		return f
	}

	return Function{
		BoundVariable: f.BoundVariable,
		Expression:    expr,
	}
}

func substituteName(n Name, before Name, after Expression) Expression {
	if n == before {
		return after
	}
	return n
}
