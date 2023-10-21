package evaluator

import "github.com/Mostafa-DE/delang/object"

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
