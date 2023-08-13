package evaluator

import "github.com/Mostafa-DE/delang/object"

func evalFunction(fun object.Object, args []object.Object) object.Object {
	function, ok := fun.(*object.Function)

	if !ok {
		return throwError("not a function: %s", fun.Type())
	}

	extendedEnv := createLocalEnv(function, args)

	evaluated := Eval(function.Body, extendedEnv)

	return unwrapReturnValue(evaluated)

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
