package evaluator

import "github.com/Mostafa-DE/delang/object"

func evalArrayIndex(array object.Object, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return throwError("Index out of range")
	}

	return arrayObject.Elements[idx]
}
