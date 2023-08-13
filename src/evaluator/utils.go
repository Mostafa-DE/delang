package evaluator

import (
	"fmt"

	"github.com/Mostafa-DE/delang/object"
)

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}

func throwError(format string, a ...interface{}) *object.Error {
	return &object.Error{Msg: fmt.Sprintf(format, a...)}
}

func getBooleanObject(val bool) *object.Boolean {
	if val {
		return TRUE
	}

	return FALSE
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
