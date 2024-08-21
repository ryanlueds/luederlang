package evaluator

import (
    "luederlang/object"
    "luederlang/ast"
    "fmt"
)

var (
    TRUE = &object.Boolean{Value: true} // what do you mean these arent const
    FALSE = &object.Boolean{Value: false}
    NULL = &object.Null{}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    case *ast.FunctionLiteral:
        params := node.Parameters
        body := node.Body
        return &object.Function{Parameters: params, Body: body, Env: env}

    case *ast.Identifier:
        return evalIdentifier(node, env)

    case *ast.CallExpression:
        function := Eval(node.Function, env)
        if isError(function) {
            return function
        }

        args := evalExpressions(node.Arguments, env)
        if len(args) == 1 && isError(args[0]) {
            return args[0]
        }

    case *ast.LetStatement:
        val := Eval(node.Value, env)
        if isError(val) {
            return val
        }
        env.Set(node.Name.Value, val)

    case *ast.IntStatement:
        val := Eval(node.Value, env)
        if isError(val) {
            return val
        }
        env.Set(node.Name.Value, val)

    case *ast.FloatStatement:
        val := Eval(node.Value, env)
        if isError(val) {
            return val
        }
        env.Set(node.Name.Value, val)

    case *ast.Program:
        return evalProgram(node, env)

    case *ast.ExpressionStatement:
        return Eval(node.Expression, env)

    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}

    case *ast.FloatLiteral:
        return &object.Float{Value: node.Value}

    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)

    case *ast.PrefixExpression:
        right := Eval(node.Right, env)
        if isError(right) {
            return right
        }
        return evalPrefixExpression(node.Operator, right)

    case *ast.InfixExpression:
        left := Eval(node.Left, env)
        if isError(left) {
            return left
        }

        right := Eval(node.Right, env)
        if isError(right) {
            return right
        }
        return evalInfixExpression(left, node.Operator, right)

    case *ast.BlockStatement:
        return evalBlockStatement(node, env)

    case *ast.IfExpression:
        return evalIfExpression(node, env)

    case *ast.ReturnStatement:
        val := Eval(node.ReturnValue, env)
        if isError(val) {
            return val
        }
        return &object.ReturnValue{Value: val}
    }
    return NULL
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
    var result []object.Object

    for _, e := range exps {
        evaluated := Eval(e, env)
        if isError(evaluated) {
            return []object.Object{evaluated}
        }
        result = append(result, evaluated)
    }

    return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
    val, ok := env.Get(node.Value)
    if !ok {
        return newError("identifier not found: %s", node.Value)
    }
    return val
}

func newError(format string, a ...interface{}) *object.Error {
    return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
    var result object.Object

    for _, statement := range program.Statements {
        result = Eval(statement, env)

        switch result := result.(type) {
        case *object.ReturnValue:
            return result.Value
        case *object.Error:
            return result
        }
    }
    return result
}

func evalBlockStatement(bs *ast.BlockStatement, env *object.Environment) object.Object {
    var result object.Object

    for _, statement := range bs.Statements {
        result = Eval(statement, env)

        if result != nil {
            rt := result.Type()
            if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
                return result
            }
        }
    }
    return result
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
    condition := Eval(ie.Condition, env)
    if isError(condition) {
        return condition
    }
    if isTruthy(condition) {
        return Eval(ie.Consequence, env)
    } else if ie.Alternative != nil {
        return Eval(ie.Alternative, env)
    }
    return NULL
}

func isTruthy(obj object.Object) bool {
    switch obj {
    case TRUE:
        return true
    case FALSE:
        return false
    case NULL:
        return false
    default:
        // while (1) {} is cool?
        return true // TODO: error handling when?
    }
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
    switch operator {
    case "!":
        return evalBangPrefixExpression(right)
    case "-":
        return evalMinusPrefixExpression(right)
    default:
        return newError("unknown operator: %s%s", operator, right.Type())
    }
}

func evalBangPrefixExpression(right object.Object) object.Object {
    switch right {
    case TRUE:
        return FALSE
    case FALSE:
        return TRUE
    case NULL:
        return TRUE
    default:
        return newError("type mismatch: !%s", right.Type())
    }
}

func evalMinusPrefixExpression(right object.Object) object.Object {
    switch right.Type() {
    case object.INTEGER_OBJ:
        return &object.Integer{Value: 0-right.(*object.Integer).Value}
    case object.FLOAT_OBJ:
        return &object.Float{Value: 0-right.(*object.Float).Value}
    default:
        return newError("type mismatch: -%s", right.Type())
    }
}

/*
 * Left and right sides don't need to be of the same type
*/
func evalInfixExpression(
    left object.Object,
    operator string,
    right object.Object,
) object.Object {
    switch operator{
    case "+":
        return evalPlusInfixExpression(left, right)
    case "*":
        return evalMultiplyInfixExpression(left, right)
    case "-":
        return evalSubtractInfixExpression(left, right)
    case "/":
        return evalDivideInfixExpression(left, right)
    case ">":
        return evalGTInfixExpression(left, right)
    case "<":
        return evalLTInfixExpression(left, right)
    case "==":
        return evalEqualsInfixExpression(left, right)
    case "!=":
        return evalNotEqualsInfixExpression(left, right)
    default:
        fmt.Printf("%s = %s", left.Type(), right.Type())
        return newError("How did you even do this... What operator is %s?", operator)
    }
}

// TODO: Figure out a better way to do this PLEASE
func evalPlusInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Integer{Value: leftVal + rightVal}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: float64(leftVal) + rightVal}
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Float{Value: leftVal + float64(rightVal)}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: leftVal + rightVal}
        }
    }
    return newError("type mismatch: %s + %s", left.Type(), right.Type())
}

func evalMultiplyInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Integer{Value: leftVal * rightVal}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: float64(leftVal) * rightVal}
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Float{Value: leftVal * float64(rightVal)}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: leftVal * rightVal}
        }
    }
    return newError("type mismatch: %s * %s", left.Type(), right.Type())
}

func evalSubtractInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Integer{Value: leftVal - rightVal}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: float64(leftVal) - rightVal}
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Float{Value: leftVal - float64(rightVal)}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: leftVal - rightVal}
        }
    }
    return newError("type mismatch: %s - %s", left.Type(), right.Type())
}

func evalDivideInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Integer{Value: leftVal / rightVal}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: float64(leftVal) / rightVal}
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return &object.Float{Value: leftVal / float64(rightVal)}
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return &object.Float{Value: leftVal / rightVal}
        }
    }
    return newError("type mismatch: %s / %s", left.Type(), right.Type())
}

func evalLTInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal < rightVal)
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(float64(leftVal) < rightVal)
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal < float64(rightVal))
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(leftVal < rightVal)
        }
    }
    return newError("type mismatch: %s < %s", left.Type(), right.Type())
}

func evalGTInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal > rightVal)
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(float64(leftVal) > rightVal)
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal > float64(rightVal))
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(leftVal > rightVal)
        }
    }
    return newError("type mismatch: %s > %s", left.Type(), right.Type())
}

func evalEqualsInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal == rightVal)
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(float64(leftVal) == rightVal)
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal == float64(rightVal))
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(leftVal == rightVal)
        }
    }
    if left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ {
        return nativeBoolToBooleanObject(left == right)
    }
    return newError("type mismatch: %s == %s", left.Type(), right.Type())
}

func evalNotEqualsInfixExpression(left, right object.Object) object.Object {
    switch left.Type() {
    case object.INTEGER_OBJ:
        leftVal := left.(*object.Integer).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal != rightVal)
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(float64(leftVal) != rightVal)
        }
    case object.FLOAT_OBJ:
        leftVal := left.(*object.Float).Value
        switch right.Type() {
        case object.INTEGER_OBJ:
            rightVal := right.(*object.Integer).Value
            return nativeBoolToBooleanObject(leftVal != float64(rightVal))
        case object.FLOAT_OBJ:
            rightVal := right.(*object.Float).Value
            return nativeBoolToBooleanObject(leftVal != rightVal)
        }
    }
    if left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ {
        return nativeBoolToBooleanObject(left != right)
    }
    return newError("type mismatch: %s != %s", left.Type(), right.Type())
}

