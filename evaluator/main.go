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

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program: // Root node of every AST our parser produces
		return evalProgram(node.Statements, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.Integer:
		return &object.Integer{Value: node.Value}

	case *ast.Float:
		return &object.Float{Value: node.Value}

	case *ast.Boolean:
		// This is because we don't need to create a new object for every boolean literal
		// Comparing the pointer address is enough
		return getBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)

		if isError(right) {
			return right
		}

		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)

		if isError(val) {
			// We need to return the error object because we need to propagate the error
			return val
		}

		return &object.Return{Value: val}

	case *ast.LetStatement:
		val := Eval(node.Value, env)

		if isError(val) {
			return val
		}

		returnValue := env.Set(node.Name.Value, val, false)

		if isError(returnValue) {
			return returnValue
		}

	case *ast.ConstStatement:
		val := Eval(node.Value, env)

		if isError(val) {
			return val
		}

		returnValue := env.Set(node.Name.Value, val, true)

		if isError(returnValue) {
			return returnValue
		}

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.Function:
		params := node.Parameters
		body := node.Body

		return &object.Function{Parameters: params, Body: body, Env: env}

	case *ast.CallFunction:
		function := Eval(node.Function, env)

		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)

		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return evalFunction(function, args, env)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.Array:
		elements := evalExpressions(node.Elements, env)

		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}

		return &object.Array{Elements: elements}

	case *ast.IndexExpression:
		ident := Eval(node.Ident, env)

		if isError(ident) {
			return ident
		}

		idx := Eval(node.Index, env)

		if isError(idx) {
			return idx
		}

		if node.Value != nil {
			val := Eval(node.Value, env)

			if isError(val) {
				return val
			}

			return setIndexExpression(ident, idx, val)
		} else {
			return evalIndexExpression(ident, idx)
		}

	case *ast.Hash:
		return evalHash(node, env)

	case *ast.AssignExpression:
		eval := evalAssignExpression(node, env)

		if isError(eval) {
			return eval
		}

		return eval

	case *ast.DuringExpression:
		return evalDuringExpression(node, env)

	case *ast.ForStatement:
		return evalForStatement(node, env)

	case *ast.BreakStatement:
		return &object.Break{}

	case *ast.SkipStatement:
		return &object.Skip{}

	}

	return nil
}

func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, env)

		if returnValue, ok := result.(*object.Return); ok {
			return returnValue.Value
		}

		if err, ok := result.(*object.Error); ok {
			return err
		}
	}

	return result
}

func evalBlockStatement(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object = &object.Null{}

	for _, statement := range statements {
		result = Eval(statement, env)

		if err, ok := result.(*object.Error); ok {
			return err
		}

		if result != nil {
			resultType := result.Type()

			if resultType == object.RETURN_OBJ ||
				resultType == object.ERROR_OBJ ||
				resultType == object.BREAK_OBJ ||
				resultType == object.SKIP_OBJ {
				return result
			}
		}
	}

	return result
}

func evalExpressions(expressions []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range expressions {
		evaluated := Eval(e, env)

		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}
