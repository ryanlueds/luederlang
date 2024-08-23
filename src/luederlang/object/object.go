package object

import (
    "fmt"
    "bytes"
    "strings"
    "luederlang/ast"
)

type ObjectType string

type Object interface {
    Type() ObjectType
    Inspect() string
}

const (
    INTEGER_OBJ = "INTEGER"
    BOOLEAN_OBJ = "BOOLEAN"
    FLOAT_OBJ = "FLOAT"
    NULL_OBJ = "NULL"
    STRING_OBJ = "STRING"
    LIST_OBJ = "LIST"
    RETURN_VALUE_OBJ = "RETURN_VALUE"
    ERROR_OBJ = "ERROR"
    FUNCTION_OBJ = "FUNCTION"
    BUILTIN_OBJ = "BUILTIN"
)

type Integer struct {
    Value int64
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type String struct {
    Value string
}

func (s *String) Inspect() string { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }

type Float struct {
    Value float64
}

func (f *Float) Inspect() string { return fmt.Sprintf("%v", f.Value) }
func (f *Float) Type() ObjectType { return FLOAT_OBJ }

type Boolean struct {
    Value bool
}

func (b *Boolean) Inspect() string { return fmt.Sprintf("%v", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct {}

func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

type ReturnValue struct {
    Value Object
}

func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Error struct {
    Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
    Parameters []*ast.Identifier
    Body *ast.BlockStatement
    Env *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fun")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
    Function BuiltinFunction
}

func (bn *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (bi *Builtin) Inspect() string { return "built in function" }





