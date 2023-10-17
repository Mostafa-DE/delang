package evaluator

import (
	"math"

	"github.com/Mostafa-DE/delang/object"
)

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)

	case left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ:
		right = &object.String{Value: right.Inspect()}
		return evalStringInfixExpression(operator, left, right)

	case left.Type() == object.INTEGER_OBJ && right.Type() == object.STRING_OBJ:
		left = &object.String{Value: left.Inspect()}
		return evalStringInfixExpression(operator, left, right)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		right = &object.Float{Value: float64(right.(*object.Integer).Value)}
		return evalFloatInfixExpression(operator, left, right)

	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		left = &object.Float{Value: float64(left.(*object.Integer).Value)}
		return evalFloatInfixExpression(operator, left, right)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.STRING_OBJ:
		left = &object.String{Value: left.Inspect()}
		return evalStringInfixExpression(operator, left, right)

	case left.Type() == object.STRING_OBJ && right.Type() == object.FLOAT_OBJ:
		right = &object.String{Value: right.Inspect()}
		return evalStringInfixExpression(operator, left, right)

	// This is pointer comparison because we only have one instance of TRUE and FALSE in memory
	// This not the case for integers because we create a new object for every integer literal
	// So we need to unwrap the object and compare the values, otherwise we would be comparing pointers
	// and that would always return false or true
	case operator == "==":
		return getBooleanObject(left == right)

	case operator == "!=":
		return getBooleanObject(left != right)

	case left.Type() != right.Type():
		return throwError("type mismatch: %s %s %s", left.Type(), operator, right.Type())

	default:
		return throwError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}

	case "-":
		return &object.Integer{Value: leftVal - rightVal}

	case "*":
		return &object.Integer{Value: leftVal * rightVal}

	case "/":
		return &object.Integer{Value: leftVal / rightVal}

	case "%":
		return &object.Integer{Value: leftVal % rightVal}

	case "<":
		return getBooleanObject(leftVal < rightVal)

	case ">":
		return getBooleanObject(leftVal > rightVal)

	case "==":
		return getBooleanObject(leftVal == rightVal)

	case "!=":
		return getBooleanObject(leftVal != rightVal)

	default:
		return throwError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}

	case "==":
		return getBooleanObject(leftVal == rightVal)

	case "!=":
		return getBooleanObject(leftVal != rightVal)

	default:
		return throwError("unknown operator: %s %s %s", left.Type(), operator, right.Type())

	}
}

func evalFloatInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}

	case "-":
		return &object.Float{Value: leftVal - rightVal}

	case "*":
		return &object.Float{Value: leftVal * rightVal}

	case "/":
		return &object.Float{Value: leftVal / rightVal}

	case "%":
		return &object.Float{Value: math.Mod(float64(leftVal), float64(rightVal))}

	case "<":
		return getBooleanObject(leftVal < rightVal)

	case ">":
		return getBooleanObject(leftVal > rightVal)

	case "==":
		return getBooleanObject(leftVal == rightVal)

	case "!=":
		return getBooleanObject(leftVal != rightVal)

	default:
		return throwError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
