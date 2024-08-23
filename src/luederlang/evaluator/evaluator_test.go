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
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64, test string) bool {
	res, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("%s || object is not an Integer. got=%T (+%v)", test, obj, obj)
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
		input    string
		expected int64
	}{
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
		testIntegerObject(t, evaluated, tt.expected, tt.input)
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{"if (10 > 1) { return 10; }", 10},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return 10;
  }

  return 1;
}
`,
			10,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected, tt.input)
	}
}

func testFloatObject(t *testing.T, obj object.Object, expected float64, test string) bool {
	res, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("%s | object is not an Float. got=%T (+%v)", test, obj, obj)
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
		input    string
		expected float64
	}{
		{"5.4", 5.4},
		{"10.051", 10.051},
		{"69.420", 69.420},
		{"-69.420", -69.420},
		{"3.5 * (3 * 3) + 10", 41.5},
		{"(5 + 10 * 2 + 15 / 3) * 2.5 + -10", 65},
		{"420.69 + 1", 421.69},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected, tt.input)
	}
}

func testNullObject(t *testing.T, obj object.Object, test string) bool {
	_, ok := obj.(*object.Null)
	if !ok {
		t.Errorf("%s | object is not a Null. got=%T (+%v)", test, obj, obj)
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
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!!true", false},
	}

	for _, tt := range test {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected, tt.input)
	}
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

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer), tt.input)
		} else {
			testNullObject(t, evaluated, tt.input)
		}
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"type mismatch: -BOOLEAN",
		},
		{
			"true + false;",
			"type mismatch: BOOLEAN + BOOLEAN",
		},
		{
			"true + false + true + false;",
			"type mismatch: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"type mismatch: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"type mismatch: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
  if (10 > 1) {
    return true + false;
  }

  return 1;
}
`,
			"type mismatch: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf(tt.input)
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf(tt.input)
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
        {"let a = 9; a = 7; a", 7},
        {"float a = 9; a = 7; a", 7}, // TODO: this shouldn't work... unless...
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected, tt.input)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fun(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fun(x) { x; }; identity(5);", 5},
		{"let identity = fun(x) { return x; }; identity(5);", 5},
		{"let double = fun(x) { x * 2; }; double(5);", 10},
		{"let add = fun(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fun(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fun(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected, tt.input)
	}
}

func TestEnclosingEnvironments(t *testing.T) {
	input := `
let first = 10;
let second = 10;
let third = 10;

let ourFunction = fun(first) {
  let second = 20;

  first + second + third;
};

ourFunction(20) + first + second;`

	testIntegerObject(t, testEval(input), 70, "It's one test")
}

func TestStringLiteral(t *testing.T) {
    input := `"Hello World!"`

    evaluated := testEval(input)
    str, ok := evaluated.(*object.String)
    if !ok {
        t.Fatalf("string is not a string")
    }

    if str.Value != "Hello World!" {
        t.Errorf("not hello world")
    }
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}
