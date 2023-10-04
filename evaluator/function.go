package evaluator

import (
	"bytes"

	"github.com/Mostafa-DE/delang/object"
)

func evalFunction(fun object.Object, args []object.Object, env *object.Environment) object.Object {
	switch fun := fun.(type) {
	case *object.Function:
		localEnv := createLocalEnv(fun, args)
		evaluated := Eval(fun.Body, localEnv)

		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		// TODO: This should be handled in a better way
		// TODO: consider moving this to logs builtin function
		if fun.Name == "logs" {
			var buffer []bytes.Buffer

			for _, arg := range args {
				buffer = append(buffer, bytes.Buffer{})
				buffer[len(buffer)-1].WriteString(arg.Inspect())
			}

			if env.GetOuterEnv() != nil {
				env = env.GetMainEnv()
			}

			if logs, ok := env.Get("bufferLogs"); ok {
				logs.(*object.Buffer).Value = append(logs.(*object.Buffer).Value, buffer...)
			} else {
				env.Set("bufferLogs", &object.Buffer{Value: buffer}, false)
			}

		}

		return fun.Func(args...)

	default:
		return throwError("not a function: %s", fun.Type())
	}

}

func createLocalEnv(fun *object.Function, args []object.Object) *object.Environment {
	env := object.NewLocalEnvironment(fun.Env)

	for idx, param := range fun.Parameters {
		env.Set(param.Value, args[idx], false)
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	}

	return obj
}
