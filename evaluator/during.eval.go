package evaluator

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalDuringExpression(node *ast.DuringExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)

	for isTruthy(condition) {
		result := evalBlockStatement(node.Body.Statements, env)

		if isError(result) {
			return result
		}

		if result.Type() == object.BREAK_OBJ {
			break
		}

		if result.Type() == object.SKIP_OBJ {
			condition = Eval(node.Condition, env)
			continue
		}

		condition = Eval(node.Condition, env)

		if isTruthy(condition) {
			continue
		} else {
			break
		}
	}

	return NULL
}
