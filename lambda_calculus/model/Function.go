package lambda_calculus

type Function struct {
	BoundVariable Name
	Expression    Expression
}

func NewFunction(boundVariable Name, expression Expression) Function {
	return Function{
		BoundVariable: boundVariable,
		Expression:    expression,
	}
}

func ReuseFunction(f Function, boundVariable Name, expression Expression) Function {
	if f.BoundVariable == boundVariable && f.Expression == expression {
		return f
	}
	return NewFunction(boundVariable, expression)
}

// Function interface

func (f Function) Accept(v ExpressionVisitor) {
	v.VisitFunction(f)
}
