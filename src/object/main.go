package object

import (
	"bytes"
	"fmt"

	"github.com/Mostafa-DE/delang/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type Return struct {
	Value Object
}

type Error struct {
	Msg string
}

type Null struct{}

type String struct {
	Value string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	// The environment in which the function was defined, This allow a closure
	Env *Environment
}

const (
	INTEGER_OBJ  = "INTEGER"
	BOOLEAN_OBJ  = "BOOLEAN"
	RETURN_OBJ   = "RETURN"
	ERROR_OBJ    = "ERROR"
	NULL_OBJ     = "NULL"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ   = "STRING"
)

func (integer *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (integer *Integer) Inspect() string {
	return fmt.Sprintf("%d", integer.Value)
}

func (boolean *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (boolean *Boolean) Inspect() string {
	return fmt.Sprintf("%t", boolean.Value)
}

func (null *Null) Type() ObjectType {
	return NULL_OBJ
}

func (null *Null) Inspect() string {
	return "null"
}

func (returnObj *Return) Type() ObjectType {
	return RETURN_OBJ
}

func (returnObj *Return) Inspect() string {
	return returnObj.Value.Inspect()
}

func (err *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (err *Error) Inspect() string {
	return "ERROR: " + err.Msg
}

func (function *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (function *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, param := range function.Parameters {
		params = append(params, param.String())
	}

	out.WriteString("fun")
	out.WriteString("(")
	out.WriteString(fmt.Sprintf("%s", params))
	out.WriteString(") {\n")
	out.WriteString(function.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (string *String) Type() ObjectType {
	return STRING_OBJ
}

func (string *String) Inspect() string {
	return string.Value
}
