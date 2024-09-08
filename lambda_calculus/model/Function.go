package lambda_calculus

type Function struct {
	BoundVariable Name
	Expression    Expression
}

// Function interface

func (f Function) Accept(v ExpressionVisitor) {
	v.VisitFunction(f)
}
