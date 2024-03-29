package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/Mostafa-DE/delang/ast"
	"github.com/shopspring/decimal"
)

type Object interface {
	Type() string
	Inspect() string
}

type Integer struct {
	Value int64
}

type Float struct {
	Value float64
}

type Decimal struct {
	Value decimal.Decimal
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

type Builtin struct {
	Func func(args ...Object) Object
	Desc string
	Name string
}

type Buffer struct {
	Value []bytes.Buffer
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	// The environment in which the function was defined, This allow a closure
	Env *Environment
}

type Array struct {
	Elements []Object
}

type HashKey struct {
	Type  string
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

type Hashable interface {
	HashKey() HashKey
}

type Break struct{}

type Skip struct{}

const (
	INTEGER_OBJ  = "INTEGER"
	FLOAT_OBJ    = "FLOAT"
	DECIMAL_OBJ  = "DECIMAL"
	BOOLEAN_OBJ  = "BOOLEAN"
	RETURN_OBJ   = "RETURN"
	ERROR_OBJ    = "ERROR"
	NULL_OBJ     = "NULL"
	FUNCTION_OBJ = "FUNCTION"
	STRING_OBJ   = "STRING"
	BUILTIN_OBJ  = "BUILTIN"
	ARRAY_OBJ    = "ARRAY"
	HASH_OBJ     = "HASH"
	BREAK_OBJ    = "BREAK"
	SKIP_OBJ     = "SKIP"
)

func (integer *Integer) Type() string {
	return INTEGER_OBJ
}

func (integer *Integer) Inspect() string {
	return fmt.Sprintf("%d", integer.Value)
}

func (float *Float) Type() string {
	return FLOAT_OBJ
}

func (float *Float) Inspect() string {
	return fmt.Sprintf("%v", float.Value)
}

func (decimal *Decimal) Type() string {
	return DECIMAL_OBJ
}

func (decimal *Decimal) Inspect() string {
	return fmt.Sprintf("%v", decimal.Value)
}

func (boolean *Boolean) Type() string {
	return BOOLEAN_OBJ
}

func (boolean *Boolean) Inspect() string {
	return fmt.Sprintf("%t", boolean.Value)
}

func (null *Null) Type() string {
	return NULL_OBJ
}

func (null *Null) Inspect() string {
	return "null"
}

func (returnObj *Return) Type() string {
	return RETURN_OBJ
}

func (returnObj *Return) Inspect() string {
	return returnObj.Value.Inspect()
}

func (err *Error) Type() string {
	return ERROR_OBJ
}

func (err *Error) Inspect() string {
	return "ERROR: " + err.Msg
}

func (function *Function) Type() string {
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
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(function.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (str *String) Type() string {
	return STRING_OBJ
}

func (str *String) Inspect() string {
	return str.Value
}

func (builtin *Builtin) Type() string {
	return BUILTIN_OBJ
}

func (builtin *Builtin) Inspect() string {
	return "builtin function" // TODO: add the name of the function
}

func (array *Array) Type() string {
	return ARRAY_OBJ
}

func (array *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, element := range array.Elements {
		elements = append(elements, element.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (buffer *Buffer) Type() string {
	return "BUFFER"
}

func (buffer *Buffer) Inspect() string {
	var out bytes.Buffer

	for _, buffer := range buffer.Value {
		out.WriteString(buffer.String())
		out.WriteString("\n")
	}

	return out.String()
}

func (hash *Hash) Type() string {
	return HASH_OBJ
}

func (hash *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}

	for _, pair := range hash.Pairs {
		pairs = append(pairs, fmt.Sprintf("'%s': '%s'", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (str *String) HashKey() HashKey {
	// TODO: Consider improve the performance by caching the hash key
	h := fnv.New64a()

	h.Write([]byte(str.Value))

	return HashKey{Type: str.Type(), Value: h.Sum64()}
}

func (b *Break) Type() string {
	return BREAK_OBJ
}

func (b *Break) Inspect() string {
	return "break"
}

func (s *Skip) Type() string {
	return SKIP_OBJ
}

func (s *Skip) Inspect() string {
	return "skip"
}
