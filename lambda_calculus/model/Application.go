package lambda_calculus

type Application struct {
	Function Expression
	Argument Expression
}

func NewApplication(function Expression, argument Expression) Application {
	return Application{Function: function, Argument: argument}
}

func ReuseApplication(a Application, function Expression, argument Expression) Application {
	if a.Function == function && a.Argument == argument {
		return a
	}
	return NewApplication(function, argument)
}

// Application interface

func (a Application) Accept(v ExpressionVisitor) {
	v.VisitApplication(a)
}
