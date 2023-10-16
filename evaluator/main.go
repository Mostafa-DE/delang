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
	var result object.Object = &object.Null{}

	for _, statement := range statements {
		result = Eval(statement, env)

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

func setIndexExpression(ident object.Object, index object.Object, value object.Object) object.Object {
	switch {
	case ident.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return setArrayIndex(ident, index, value)

	case ident.Type() == object.HASH_OBJ:
		return setHashIndex(ident, index, value)

	default:
		return throwError("index operator not supported: %s", ident.Type())

	}
}

func setArrayIndex(array object.Object, index object.Object, value object.Object) object.Object {
	arrayObject := array.(*object.Array)

	idx := index.(*object.Integer).Value

	if idx < 0 || idx > int64(len(arrayObject.Elements)-1) {
		return throwError("index out of bounds")
	}

	arrayObject.Elements[idx] = value

	return NULL
}

func setHashIndex(hash object.Object, index object.Object, value object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)

	if !ok {
		return throwError("Type %s is not hashable", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]

	if !ok {
		return throwError("Index not found")
	}

	pair.Value = value

	hashObject.Pairs[key.HashKey()] = pair

	return NULL
}

func evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
	idxIdent := node.IdxIdent.Value
	varIdent := node.VarIdent.Value
	body := node.Body

	if idxIdent == varIdent {
		return throwError("Index identifier and variable identifier cannot be the same")
	}

	if node.Expression == nil {
		return throwError("Expected an expression after for statement")
	}

	eval := Eval(node.Expression, env)

	if isError(eval) {
		return eval
	}

	iterable := eval

	switch iterable.Type() {
	case object.STRING_OBJ:
		stringLoop(iterable.(*object.String), idxIdent, varIdent, body, env)

	case object.ARRAY_OBJ:
		arrayLoop(iterable.(*object.Array), idxIdent, varIdent, body, env)

	default:
		return throwError("Type %s is not iterable", iterable.Type())
	}

	return NULL
}

func arrayLoop(array *object.Array, idxIdent string, varIdent string, body *ast.BlockStatement, env *object.Environment) {
	for idx, val := range array.Elements {
		env.Set(idxIdent, &object.Integer{Value: int64(idx)}, false)
		env.Set(varIdent, val, false)

		result := evalBlockStatement(body.Statements, env)

		if isError(result) {
			return
		}

		if result != nil {
			if result.Type() == object.BREAK_OBJ {
				break
			}

			if result.Type() == object.SKIP_OBJ {
				continue
			}
		}
	}
}

func stringLoop(str *object.String, idxIdent string, varIdent string, body *ast.BlockStatement, env *object.Environment) {
	for idx, val := range str.Value {
		env.Set(idxIdent, &object.Integer{Value: int64(idx)}, false)
		env.Set(varIdent, &object.String{Value: string(val)}, false)

		result := evalBlockStatement(body.Statements, env)

		if isError(result) {
			return
		}

		if result != nil {
			if result.Type() == object.BREAK_OBJ {
				break
			}

			if result.Type() == object.SKIP_OBJ {
				continue
			}
		}
	}
}
