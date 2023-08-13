package evaluator

import (
	"github.com/Mostafa-DE/delang/ast"
	"github.com/Mostafa-DE/delang/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	return throwError("identifier not found: %s", node.Value)
}
