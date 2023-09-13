package object

// Environment is a map of strings to Objects that we can use to store and retrieve values
// from the environment we're currently in (e.g. global or local)

type StoreType map[string]Object

type Environment struct {
	store       StoreType
	constValues map[string]struct{}
	outer       *Environment
}

func NewEnvironment() *Environment {
	s := make(StoreType)

	return &Environment{store: s, outer: nil, constValues: make(map[string]struct{})}
}

func NewLocalEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, val Object, isConst bool) Object {
	if _, ok := e.store[name]; ok {
		if isConst {
			return throwError("Cannot redeclare constant '%s'", name)
		}
	}

	if _, ok := e.constValues[name]; ok {
		return throwError("Cannot reassign constant '%s'", name)
	}

	e.store[name] = val

	if isConst {
		e.constValues[name] = struct{}{}
	}

	return val
}
