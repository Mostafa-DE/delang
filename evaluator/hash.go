package evaluator

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalHash(node *ast.Hash, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for key, val := range node.Pairs {
		key := Eval(key, env)

		if isError(key) {
			return key
		}

		// Check if the key is hashable
		hashKey, ok := key.(object.Hashable)

		if !ok {
			return throwError("Type %s is not hashable", key.Type())
		}

		value := Eval(val, env)

		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()

		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func evalHashIndex(hash object.Object, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)

	if !ok {
		return throwError("Type %s is not hashable", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]

	if !ok {
		return NULL
	}

	return pair.Value
}
