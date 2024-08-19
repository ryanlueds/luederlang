package evaluator

import (
    "luederlang/object"
    "luederlang/ast"
)

var (
    TRUE = &object.Boolean{Value: true} // what do you mean these arent const
    FALSE = &object.Boolean{Value: false}
    NULL = &object.Null{}
)

/*var typeToHash = map[object.ObjectType]int8{
    object.INTEGER_OBJ:1 << 0,
    object.FLOAT_OBJ: 1 << 1,
}

func getTypesHash(left, right object.Object) int8 {
    leftHash, ok := typeToHash[left.Type()]
    if !ok {
        // TODO: there's no way i can keep doing this everywhere
        return 0
    }
    rightHash, ok := typeToHash[right.Type()]
    if !ok {
        return 0
    }
    return leftHash + rightHash
}

func getValueFromObject[T any](obj object.Object) T {
    switch obj.Type() {
    case object.Integer:
        return obj.(object.Integer
    
    }
}*/

func Eval(node ast.Node) object.Object {
    switch node := node.(type) {
    case *ast.Program:
        return evalStatements(node.Statements)
    case *ast.ExpressionStatement:
        return Eval(node.Expression)
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    case *ast.FloatLiteral:
        return &object.Float{Value: node.Value}
    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)
    case *ast.PrefixExpression:
        return evalPrefixExpression(node.Operator, Eval(node.Right))
    case *ast.InfixExpression:
        return evalInfixExpression(Eval(node.Left), node.Operator, Eval(node.Right))
    }
    return nil
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(statements []ast.Statement) object.Object {
    var result object.Object

    for _, statement := range statements {
        result = Eval(statement)
    }
    
    return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
    switch operator {
    case "!":
        return evalBangPrefixExpression(right)
    case "-":
        return evalMinusPrefixExpression(right)
    default:
        return NULL
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
        // TODO: error handling, this should never happen
        return NULL
    }
}

func evalMinusPrefixExpression(right object.Object) object.Object {
    switch right.Type() {
    case object.INTEGER_OBJ:
        return &object.Integer{Value: 0-right.(*object.Integer).Value}
    case object.FLOAT_OBJ:
        return &object.Float{Value: 0-right.(*object.Float).Value}
    default:
        return NULL
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
        return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
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
    // TODO: Figure out error handling
    return NULL
}

