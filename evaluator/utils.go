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

	if obj == NULL {
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

func intToBool(val int64) bool {
	return val != 0
}

func makeRangeArray(startRange int64, endRange int64) *object.Array {
	size := endRange - startRange + 1
	elements := make([]object.Object, size)

	for i := int64(0); i < size; i++ {
		elements[i] = &object.Integer{Value: startRange + i}
	}

	return &object.Array{Elements: elements}
}
