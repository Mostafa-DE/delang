package evaluator

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	localEnv := object.NewLocalEnvironment(env)
	condition := Eval(ie.Condition, localEnv)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, localEnv)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, localEnv)
	} else {
		return NULL
	}
}
