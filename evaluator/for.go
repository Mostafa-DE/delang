package evaluator

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
	localEnv := object.NewLocalEnvironment(env)
	var idxIdent string
	if node.IdxIdent != nil {
		idxIdent = node.IdxIdent.Value
	}

	var varIdent string
	if node.VarIdent == nil {
		return throwError("Expected a variable identifier after for statement")
	} else {
		varIdent = node.VarIdent.Value
	}

	body := node.Body

	if idxIdent == varIdent {
		return throwError("Index identifier and variable identifier cannot be the same")
	}

	if node.Expression == nil {
		return throwError("Expected an expression after for statement")
	}

	eval := Eval(node.Expression, localEnv)

	if isError(eval) {
		return eval
	}

	switch eval.Type() {
	case object.STRING_OBJ:
		res := stringLoop(eval.(*object.String), idxIdent, varIdent, body, localEnv)

		if isError(res) {
			return res
		}

	case object.ARRAY_OBJ:
		res := arrayLoop(eval.(*object.Array), idxIdent, varIdent, body, localEnv)

		if isError(res) {
			return res
		}

	default:
		return throwError("Type %s is not iterable", eval.Type())
	}

	return NULL
}

func arrayLoop(
	array *object.Array,
	idxIdent string,
	varIdent string, body *ast.BlockStatement,
	env *object.Environment,
) object.Object {
	for idx, val := range array.Elements {
		if idxIdent != "" {
			env.Set(idxIdent, &object.Integer{Value: int64(idx)}, false)
		}

		env.Set(varIdent, val, false)

		result := evalBlockStatement(body.Statements, env)

		if isError(result) {
			return result
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

	return NULL
}

func stringLoop(
	str *object.String,
	idxIdent string,
	varIdent string,
	body *ast.BlockStatement,
	env *object.Environment,
) object.Object {
	for idx, val := range str.Value {
		env.Set(idxIdent, &object.Integer{Value: int64(idx)}, false)
		env.Set(varIdent, &object.String{Value: string(val)}, false)

		result := evalBlockStatement(body.Statements, env)

		if isError(result) {
			return result
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

	return NULL
}
