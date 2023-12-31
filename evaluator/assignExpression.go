package evaluator

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalAssignExpression(node *ast.AssignExpression, env *object.Environment) object.Object {
	val := Eval(node.Value, env)

	if isError(val) {
		return val
	}

	_, ok := env.Get(node.Ident.Value)

	if !ok {
		return throwError("identifier not found: %s", node.Ident.Value)
	}

	returnVal := env.GetTargetEnv(node.Ident.Value).Set(node.Ident.Value, val, false)

	if isError(returnVal) {
		return returnVal
	}

	return returnVal
}
