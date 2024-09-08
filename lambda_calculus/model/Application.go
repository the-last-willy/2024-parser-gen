package lambda_calculus

type Application struct {
	Function Expression
	Argument Expression
}

// Application interface

func (a Application) Accept(v ExpressionVisitor) {
	v.VisitApplication(a)
}
