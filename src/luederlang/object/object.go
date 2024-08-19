package object

import (
    "fmt"
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
)

type Integer struct {
    Value int64
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

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









