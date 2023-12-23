package evaluator

import (
	"math"

	"github.com/Mostafa-DE/delang/object"
	"github.com/shopspring/decimal"
)

func evalInfixExpression(operator string, left object.Object, right object.Object, env *object.Environment) object.Object {
	_leftType := left.Type()
	_rightType := right.Type()

	switch {
	case _leftType == object.INTEGER_OBJ && _rightType == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	case _leftType == object.STRING_OBJ && _rightType == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)

	case _leftType == object.STRING_OBJ && _rightType == object.INTEGER_OBJ:
		right = &object.String{Value: right.Inspect()}

		return evalStringInfixExpression(operator, left, right)

	case _leftType == object.INTEGER_OBJ && _rightType == object.STRING_OBJ:
		left = &object.String{Value: left.Inspect()}

		return evalStringInfixExpression(operator, left, right)

	case _leftType == object.FLOAT_OBJ && _rightType == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right)

	case _leftType == object.FLOAT_OBJ && _rightType == object.INTEGER_OBJ:
		right = &object.Float{Value: float64(right.(*object.Integer).Value)}

		return evalFloatInfixExpression(operator, left, right)

	case _leftType == object.INTEGER_OBJ && _rightType == object.FLOAT_OBJ:
		left = &object.Float{Value: float64(left.(*object.Integer).Value)}

		return evalFloatInfixExpression(operator, left, right)

	case _leftType == object.FLOAT_OBJ && _rightType == object.STRING_OBJ:
		left = &object.String{Value: left.Inspect()}

		return evalStringInfixExpression(operator, left, right)

	case _leftType == object.STRING_OBJ && _rightType == object.FLOAT_OBJ:
		right = &object.String{Value: right.Inspect()}

		return evalStringInfixExpression(operator, left, right)

	case _leftType == object.DECIMAL_OBJ && _rightType == object.DECIMAL_OBJ:
		return evalDecimalInfixExpression(operator, left, right, env)

	case _leftType == object.DECIMAL_OBJ && _rightType == object.INTEGER_OBJ:
		right = &object.Decimal{Value: decimal.NewFromInt(right.(*object.Integer).Value)}

		return evalDecimalInfixExpression(operator, left, right, env)

	case _leftType == object.INTEGER_OBJ && _rightType == object.DECIMAL_OBJ:
		left = &object.Decimal{Value: decimal.NewFromInt(left.(*object.Integer).Value)}

		return evalDecimalInfixExpression(operator, left, right, env)

	case _leftType == object.DECIMAL_OBJ && _rightType == object.FLOAT_OBJ:
		right = &object.Decimal{Value: decimal.NewFromFloat(right.(*object.Float).Value)}

		return evalDecimalInfixExpression(operator, left, right, env)

	case _leftType == object.FLOAT_OBJ && _rightType == object.DECIMAL_OBJ:
		left = &object.Decimal{Value: decimal.NewFromFloat(left.(*object.Float).Value)}

		return evalDecimalInfixExpression(operator, left, right, env)

	case _leftType == object.DECIMAL_OBJ && _rightType == object.STRING_OBJ:
		left = &object.String{Value: left.Inspect()}
		right = &object.String{Value: right.Inspect()}

		return evalStringInfixExpression(operator, left, right)

	case _leftType == object.STRING_OBJ && _rightType == object.DECIMAL_OBJ:
		left = &object.String{Value: left.Inspect()}
		right = &object.String{Value: right.Inspect()}

		return evalStringInfixExpression(operator, left, right)

	/*
		- This is pointer comparison because we only have one instance of TRUE and FALSE in memory
		- This not the case for integers because we create a new object for every integer literal
		- So we need to unwrap the object and compare the values.
		- otherwise we would be comparing pointers and that would always return false or true
	*/
	case operator == "==":
		return getBooleanObject(left == right)

	case operator == "!=":
		return getBooleanObject(left != right)

	case operator == "and":
		/*
			- If the 'and' operator is used with booleans and integers.
				- First we convert the integer to boolean
				- Then we compare the two booleans
		*/
		if _leftType == object.INTEGER_OBJ && _rightType == object.BOOLEAN_OBJ {
			left = getBooleanObject(intToBool(left.(*object.Integer).Value))

		} else if _leftType == object.BOOLEAN_OBJ && _rightType == object.INTEGER_OBJ {
			right = getBooleanObject(intToBool(right.(*object.Integer).Value))

		}

		return getBooleanObject((left == TRUE) && (right == TRUE))

	case operator == "or":
		/*
			- If the 'or' operator is used with booleans and integers.
				- First we convert the integer to boolean
				- Then we compare the two booleans
		*/
		if _leftType == object.INTEGER_OBJ && _rightType == object.BOOLEAN_OBJ {
			left = getBooleanObject(intToBool(left.(*object.Integer).Value))

		} else if _leftType == object.BOOLEAN_OBJ && _rightType == object.INTEGER_OBJ {
			right = getBooleanObject(intToBool(right.(*object.Integer).Value))

		}

		return getBooleanObject((left == TRUE) || (right == TRUE))

	case _leftType != _rightType:
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
		if rightVal == 0 {
			return throwError("division by zero")
		}

		return &object.Integer{Value: leftVal / rightVal}

	case "%":
		if rightVal == 0 {
			return throwError("division by zero")
		}

		return &object.Integer{Value: leftVal % rightVal}

	case "<":
		return getBooleanObject(leftVal < rightVal)

	case ">":
		return getBooleanObject(leftVal > rightVal)

	case "<=":
		return getBooleanObject(leftVal <= rightVal)

	case ">=":
		return getBooleanObject(leftVal >= rightVal)

	case "==":
		return getBooleanObject(leftVal == rightVal)

	case "!=":
		return getBooleanObject(leftVal != rightVal)

	case "and":
		/*
			- If 'and' operator is used with integers.
				- Check if the left or right value is 0
					- Then return 0

				- Otherwise return the right value
		*/

		if leftVal == 0 || rightVal == 0 {
			return &object.Integer{Value: 0}
		}

		return &object.Integer{Value: rightVal}

	case "or":
		/*
			- If 'or' operator is used with integers.
				- Check if the left value is 0
					- Then return the right value

				- Check if the right value is 0
					- Then return the left value

				- Otherwise return the left value
		*/

		if leftVal == 0 && rightVal != 0 {
			return &object.Integer{Value: rightVal}

		} else if leftVal != 0 && rightVal == 0 {
			return &object.Integer{Value: leftVal}

		}

		return &object.Integer{Value: leftVal}

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
		if rightVal == 0 {
			return throwError("division by zero")
		}

		return &object.Float{Value: leftVal / rightVal}

	case "%":
		if rightVal == 0 {
			return throwError("division by zero")
		}

		return &object.Float{Value: math.Mod(float64(leftVal), float64(rightVal))}

	case "<":
		return getBooleanObject(leftVal < rightVal)

	case ">":
		return getBooleanObject(leftVal > rightVal)

	case "<=":
		return getBooleanObject(leftVal <= rightVal)

	case ">=":
		return getBooleanObject(leftVal >= rightVal)

	case "==":
		return getBooleanObject(leftVal == rightVal)

	case "!=":
		return getBooleanObject(leftVal != rightVal)

	case "and":
		/*
			- If 'and' operator is used with floats.
				- Check if the left or right value is 0
					- Then return 0

				- Otherwise return the right value
		*/

		if leftVal == 0 || rightVal == 0 {
			return &object.Float{Value: 0}
		}

		return &object.Float{Value: rightVal}

	case "or":
		/*
			- If 'or' operator is used with floats.
				- Check if the left value is 0
					- Then return the right value

				- Check if the right value is 0
					- Then return the left value

				- Otherwise return the left value

		*/

		if leftVal == 0 && rightVal != 0 {
			return &object.Float{Value: rightVal}

		} else if leftVal != 0 && rightVal == 0 {
			return &object.Float{Value: leftVal}

		}

		return &object.Float{Value: leftVal}

	default:
		return throwError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalDecimalInfixExpression(operator string, left object.Object, right object.Object, env *object.Environment) object.Object {
	leftVal := left.(*object.Decimal).Value
	rightVal := right.(*object.Decimal).Value

	decimalData, ok := env.Get("_getDecimalData")

	if !ok {
		return throwError("_getDecimalData() not found")
	}

	decimalHash := decimalData.(*object.Hash)
	hashPrec := (&object.String{Value: "prec"}).HashKey()
	hashDivPrecString := (&object.String{Value: "divPrec"}).HashKey()

	precObj, precOk := decimalHash.Pairs[hashPrec]
	divisionPrecisionObj, objOk := decimalHash.Pairs[hashDivPrecString]

	if !precOk {
		return throwError("prec not found")
	}

	if !objOk {
		return throwError("divPrec not found")
	}

	prec, precOk := precObj.Value.(*object.Integer)

	divPrec, intOk := divisionPrecisionObj.Value.(*object.Integer)

	if !precOk {
		return throwError("prec is not an integer")
	}

	if !intOk {
		return throwError("divPrec is not an integer")
	}

	if prec.Value < 0 || prec.Value > 8 {
		return throwError("Valid range for prec is [0 to 8]")
	}

	if divPrec.Value < 0 || divPrec.Value > 28 {
		return throwError("Valid range for divPrec is [0 to 28]")
	}

	switch operator {
	case "+":
		return &object.Decimal{Value: leftVal.Add(rightVal).Round(int32(prec.Value))}

	case "-":
		return &object.Decimal{Value: leftVal.Sub(rightVal).Round(int32(prec.Value))}

	case "*":
		return &object.Decimal{Value: leftVal.Mul(rightVal).Round(int32(prec.Value))}

	case "/":
		if rightVal.IsZero() {
			return throwError("division by zero")
		}

		decimal.DivisionPrecision = int(divPrec.Value)

		return &object.Decimal{Value: leftVal.Div(rightVal)}

	case "%":
		if rightVal.IsZero() {
			return throwError("division by zero")
		}

		decimal.DivisionPrecision = int(divPrec.Value)

		return &object.Decimal{Value: leftVal.Mod(rightVal)}

	case "<":
		return getBooleanObject(leftVal.LessThan(rightVal))

	case ">":
		return getBooleanObject(leftVal.GreaterThan(rightVal))

	case "<=":
		return getBooleanObject(leftVal.LessThanOrEqual(rightVal))

	case ">=":
		return getBooleanObject(leftVal.GreaterThanOrEqual(rightVal))

	case "==":
		return getBooleanObject(leftVal.Equal(rightVal))

	case "!=":
		return getBooleanObject(!leftVal.Equal(rightVal))

	case "and":
		/*
			- If the 'and' operator is used with decimals.
				- Check if the left or right value is 0
					- Then return 0

				- Otherwise return the right value

		*/

		if leftVal.IsZero() || rightVal.IsZero() {
			return &object.Decimal{Value: decimal.Zero}
		}

		return &object.Decimal{Value: rightVal}

	case "or":
		/*
			- If the 'or' operator is used with decimals.
				- Check if the left value is 0
					- Then return the right value

				- Check if the right value is 0
					- Then return the left value

				- Otherwise return the left value

		*/

		if leftVal.IsZero() && !rightVal.IsZero() {
			return &object.Decimal{Value: rightVal}

		} else if !leftVal.IsZero() && rightVal.IsZero() {

			return &object.Decimal{Value: leftVal}

		}

		return &object.Decimal{Value: leftVal}

	default:
		return throwError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
