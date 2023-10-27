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
	s := map[string]Object{
		"getDecimalData": decimalData(),
	}

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
	if checkShadowing(name) {
		return throwError("Shadowing of '%s' is not allowed", name)
	}

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

func (e *Environment) GetOuterEnv() *Environment {
	return e.outer
}

func (e *Environment) GetMainEnv() *Environment {
	if e.outer != nil {
		return e.outer.GetMainEnv()
	}

	return e
}

func decimalData() *Hash {
	return &Hash{
		Pairs: map[HashKey]HashPair{
			(&String{Value: "prec"}).HashKey(): {
				Key: &String{
					Value: "prec",
				},
				Value: &Integer{
					Value: 8,
				},
			},
			(&String{Value: "divPrec"}).HashKey(): {
				Key: &String{
					Value: "divPrec",
				},
				Value: &Integer{
					Value: 28,
				},
			},
		},
	}
}

func checkShadowing(name string) bool {
	arr := []string{
		"getDecimalData",
		"len",
		"first",
		"last",
		"skipFirst",
		"skipLast",
		"push",
		"pop",
		"logs",
		"range",
		"decimal",
		"typeof",
	}

	for _, v := range arr {
		if name == v {
			return true
		}
	}

	return false
}
