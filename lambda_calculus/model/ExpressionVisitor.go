package lambda_calculus

type ExpressionVisitor interface {
	VisitApplication(Application)
	VisitFunction(Function)
	VisitName(Name)
}
