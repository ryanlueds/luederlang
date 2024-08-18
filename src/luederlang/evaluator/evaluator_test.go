package evaluator

import (
    "luederlang/lexer"
    "luederlang/object"
    "luederlang/parser"
    "testing"
)

func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()
    
    return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
    res, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("object is not an Integer. got=%T (+%v)", obj, obj)
        return false
    }
    if res.Value != expected {
        t.Errorf("object has wrong value. got=%d. want=%d", res.Value, expected)
        return false
    }
    return true
}

func TestEvalIntegerExpression(t *testing.T) {
    tests := []struct {
        input string
        expected int64
    } {
        {"5", 5},
        {"10", 10},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
    res, ok := obj.(*object.Float)
    if !ok {
        t.Errorf("object is not an Float. got=%T (+%v)", obj, obj)
        return false
    }
    if res.Value != expected {
        t.Errorf("object has wrong value. got=%v. want=%v", res.Value, expected)
        return false
    }
    return true
}


func TestEvalFloatExpression(t *testing.T) {
    tests := []struct {
        input string
        expected float64
    } {
        {"5.4", 5.4},
        {"10.051", 10.051},
        {"69.420", 69.420},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testFloatObject(t, evaluated, tt.expected)
    }
}
