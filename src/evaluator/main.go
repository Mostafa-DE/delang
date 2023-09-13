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

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

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

		return evalIndexExpression(ident, idx)

	case *ast.Hash:
		return evalHash(node, env)

	case *ast.AssignExpression:
		eval := evalAssignExpression(node, env)

		if isError(eval) {
			return eval
		}

		return eval

	}

	return nil
}

func evalAssignExpression(node *ast.AssignExpression, env *object.Environment) object.Object {
	val := Eval(node.Value, env)

	if isError(val) {
		return val
	}

	_, ok := env.Get(node.Ident.Value)

	if !ok {
		return throwError("identifier not found: %s", node.Ident.Value)
	}

	envVal := env.Set(node.Ident.Value, val, false)

	if isError(envVal) {
		return envVal
	}

	return val
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
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, env)

		if result != nil {
			resultType := result.Type()

			if resultType == object.RETURN_OBJ || resultType == object.ERROR_OBJ {
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

func evalIndexExpression(ident object.Object, index object.Object) object.Object {
	switch {
	case ident.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndex(ident, index)

	case ident.Type() == object.HASH_OBJ:
		return evalHashIndex(ident, index)

	default:
		return throwError("index operator not supported: %s", ident.Type())

	}
}

func evalArrayIndex(array object.Object, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		// TODO: For now we return NULL, but we need to return an error
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalHashIndex(hash object.Object, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)

	if !ok {
		return throwError("Type %s is not hashable", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]

	if !ok {
		return NULL
	}

	return pair.Value
}

func evalHash(node *ast.Hash, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for key, val := range node.Pairs {
		key := Eval(key, env)

		if isError(key) {
			return key
		}

		// Check if the key is hashable
		hashKey, ok := key.(object.Hashable)

		if !ok {
			return throwError("Type %s is not hashable", key.Type())
		}

		value := Eval(val, env)

		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()

		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}
