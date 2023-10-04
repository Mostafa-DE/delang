package evaluator

import (
	"fmt"
	"time"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalDuringExpression(node *ast.DuringExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	timeoutLoop, _ := env.Get("timeoutLoop")
	var timeout <-chan time.Time

	if timeoutLoop != nil {
		timeout = time.After(5 * time.Second)
	}

loop:
	for isTruthy(condition) {

		select {
		case <-timeout:
			fmt.Println("Timeout exceeded")
			env.Set("timeoutExceeded", &object.Boolean{Value: true}, false)
			break loop

		default:
			result := evalBlockStatement(node.Body.Statements, env)

			if isError(result) {
				return result
			}

			if result != nil {
				if result.Type() == object.BREAK_OBJ {
					break loop
				}

				if result.Type() == object.SKIP_OBJ {
					condition = Eval(node.Condition, env)
					continue loop
				}
			}

			condition = Eval(node.Condition, env)

			if isTruthy(condition) {
				continue loop
			} else {
				break loop
			}
		}
	}

	return NULL

}
