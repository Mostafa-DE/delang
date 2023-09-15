package evaluator

import "github.com/Mostafa-DE/delang/object"

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalExclamationOperatorExpression(right)

	case "-":
		return evalMinusPrefixOperatorExpression(right)

	default:
		return throwError("unknown operator: %s%s", operator, right.Type())
	}
}
