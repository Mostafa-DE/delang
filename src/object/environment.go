package object

// Environment is a map of strings to Objects that we can use to store and retrieve values
// from the environment we're currently in (e.g. global or local)

type StoreType map[string]Object

type Environment struct {
	store StoreType
}

func NewEnvironment() *Environment {
	s := make(StoreType)
	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
