package lambda_calculus

type Expression interface {
	Accept(ExpressionVisitor)
}
