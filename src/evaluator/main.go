package evaluator

import (
	"fmt"

	"github.com/Mostafa-DE/delang/object"

	"github.com/Mostafa-DE/delang/ast"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program: // Root node of every AST our parser produces
		return evalProgram(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		// This is because we don't need to create a new object for every boolean literal
		// We can just return the same object
		return getBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)

		if isError(right) {
			return right
		}

		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left) {
			return left
		}

		right := Eval(node.Right)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements)

	case *ast.IfExpression:
		return evalIfExpression(node)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)

		if isError(val) {
			// We need to return the error object because we need to propagate the error
			return val
		}

		return &object.Return{Value: val}

	}

	return nil
}

func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)

		if returnValue, ok := result.(*object.Return); ok {
			return returnValue.Value
		}

		if err, ok := result.(*object.Error); ok {
			return err
		}
	}

	return result
}

func evalBlockStatement(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)

		if result != nil {
			resultType := result.Type()

			if resultType == object.RETURN_OBJ || resultType == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

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

func evalInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

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

func evalExclamationOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE

	case FALSE:
		return TRUE

	case NULL:
		return TRUE

	default:
		return FALSE

	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return throwError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
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

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	if obj == nil {
		return false
	}

	if obj.Type() == object.INTEGER_OBJ {
		return obj.(*object.Integer).Value != 0
	}

	if obj.Type() == object.BOOLEAN_OBJ {
		return obj.(*object.Boolean).Value
	}

	return true
}

func getBooleanObject(val bool) *object.Boolean {
	if val {
		return TRUE
	}

	return FALSE
}

func throwError(format string, a ...interface{}) *object.Error {
	return &object.Error{Msg: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}
