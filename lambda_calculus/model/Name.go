package lambda_calculus

type Name string

func NewName(name string) Name {
	return Name(name)
}

// Expression interface

func (n Name) Accept(v ExpressionVisitor) {
	v.VisitName(n)
}
