package evaluator

import "github.com/Mostafa-DE/delang/object"

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
