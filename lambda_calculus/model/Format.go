package lambda_calculus

import "fmt"

func Format(e Expression) string {
	switch e := e.(type) {
	case Application:
		return formatApplication(e)
	case Function:
		return formatFunction(e)
	case Name:
		return formatName(e)
	default:
		panic("Format: unexpected Expression type")
	}
}

func formatApplication(a Application) string {
	fun := Format(a.Function)
	arg := Format(a.Argument)
	return fmt.Sprintf("(%s %s)", fun, arg)
}

func formatFunction(f Function) string {
	bv := formatName(f.BoundVariable)
	expr := Format(f.Expression)
	return fmt.Sprintf(`\%s.%s`, bv, expr)
}

func formatName(n Name) string {
	return string(n)
}
