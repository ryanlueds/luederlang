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
        {"-5", -5},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
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
        {"-69.420", -69.420},
		{"3.5 * (3 * 3) + 10", 41.5},
		{"(5 + 10 * 2 + 15 / 3) * 2.5 + -10", 65},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testFloatObject(t, evaluated, tt.expected)
    }
}

func testNullObject(t *testing.T, obj object.Object) bool {
    _, ok := obj.(*object.Null)
    if !ok {
        t.Errorf("object is not a Null. got=%T (+%v)", obj, obj)
        return false
    }
    return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool, test string) bool {
    res, ok := obj.(*object.Boolean)
    if !ok {
        t.Errorf("%s | object is not a Boolean. got=%T (+%v)", test, obj, obj)
        return false
    }
    if res.Value != expected {
        t.Errorf("object has wrong value. got=%t. want=%t", res.Value, expected)
        return false
    }
    return true
}

func TestBangOperator(t *testing.T) {
    test := []struct {
        input string
        expected bool
    } {
        {"!true", false},
        {"!false", true},
        {"!!true", true},
        {"!!!true", false},
    }

    for _, tt := range test {
        evaluated := testEval(tt.input)
        testBooleanObject(t, evaluated, tt.expected, tt.input)
    }
    
    evaluated := testEval("!x")
    testNullObject(t, evaluated)
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected, tt.input)
	}
}
