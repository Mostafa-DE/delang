package evaluator

import "github.com/Mostafa-DE/delang/object"

func evalFunction(fun object.Object, args []object.Object) object.Object {
	switch fun := fun.(type) {
	case *object.Function:
		localEnv := createLocalEnv(fun, args)
		evaluated := Eval(fun.Body, localEnv)

		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fun.Func(args...)

	default:
		return throwError("not a function: %s", fun.Type())
	}

}

func createLocalEnv(fun *object.Function, args []object.Object) *object.Environment {
	env := object.NewLocalEnvironment(fun.Env)

	for idx, param := range fun.Parameters {
		env.Set(param.Value, args[idx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	}

	return obj
}
