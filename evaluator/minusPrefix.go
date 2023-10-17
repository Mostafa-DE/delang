package evaluator

import "github.com/Mostafa-DE/delang/object"

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right.Type() {
	case object.INTEGER_OBJ:
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}

	case object.FLOAT_OBJ:
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}

	default:
		return throwError("unknown operator: -%s", right.Type())
	}
}
