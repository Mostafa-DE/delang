package evaluator

import (
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
