package evaluator

import (
	"fmt"

	"github.com/Mostafa-DE/delang/object"
)

// TODO: Revisit this file and refactor it

var builtins = map[string]*object.Builtin{
	"len": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return throwError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			default:
				return throwError("argument to `len` not supported, got %s", args[0].Type())

			}
		},
		Desc: "Returns the length of a string or an array",
		Name: "len",
	},

	"first": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return throwError("wrong number of arguments passed to first(). got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)

			if !ok {
				return throwError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return NULL
		},
		Desc: "Returns the first element of an array",
		Name: "first",
	},

	"last": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return throwError("wrong number of arguments passed to last(). got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)

			if !ok {
				return throwError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			length := len(array.Elements)

			if length > 0 {
				return array.Elements[length-1]
			}

			return NULL
		},
		Desc: "Returns the last element of an array",
		Name: "last",
	},

	"skipFirst": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return throwError("wrong number of arguments passed to skipFirst(). got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)

			if !ok {
				return throwError("argument to `skipFirst` must be ARRAY, got %s", args[0].Type())
			}

			length := len(array.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, array.Elements[1:length])

				return &object.Array{Elements: newElements}
			}

			return NULL
		},
		Desc: "Returns an array with the first element removed",
		Name: "skipFirst",
	},

	"skipLast": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return throwError("wrong number of arguments passed to skipLast(). got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)

			if !ok {
				return throwError("argument to `skipLast` must be ARRAY, got %s", args[0].Type())
			}

			length := len(array.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, array.Elements[0:length-1])

				return &object.Array{Elements: newElements}
			}

			return NULL
		},
		Desc: "Returns an array with the last element removed",
		Name: "skipLast",
	},

	"push": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return throwError("wrong number of arguments passed to push(). got=%d, want=2", len(args))
			}

			array, ok := args[0].(*object.Array)

			if !ok {
				return throwError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			array.Elements = append(array.Elements, args[1])

			return array
		},
		Desc: "Pushes an element to the end of an array",
		Name: "push",
	},

	"pop": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return throwError("wrong number of arguments passed to pop(). got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)

			if !ok {
				return throwError("argument to `pop` must be ARRAY, got %s", args[0].Type())
			}

			length := len(array.Elements)

			if length > 0 {
				array.Elements = array.Elements[0 : length-1]

				return array
			}

			return NULL
		},
		Desc: "Removes the last element of an array",
		Name: "pop",
	},

	"logs": {
		Func: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
		Desc: "Prints the result to the console",
		Name: "logs",
	},
}
