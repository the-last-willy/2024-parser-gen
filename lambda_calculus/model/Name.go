package lambda_calculus

type Name string

func NewName(s string) Name {
	return Name(s)
}

func ReuseName(n Name, s string) Name {
	if string(n) == s {
		return n
	}
	return NewName(s)
}

// Expression interface

func (n Name) Accept(v ExpressionVisitor) {
	v.VisitName(n)
}
