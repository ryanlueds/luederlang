package evaluator

import (
    "luederlang/object"
    "fmt"
)

var builtins = map[string]*object.Builtin{
    "len": &object.Builtin{
        Function: func(args ...object.Object) object.Object {
            if len(args) != 1 {
                return newError("wrong number of arguments. want=1. got=%v", len(args))
            }
            switch arg := args[0].(type) {
            case *object.String:
                return &object.Integer{Value: int64(len(arg.Value))}
            default:
                return newError("len operation only supported on strings")
            }
        },
    },
    "help": &object.Builtin{
        Function: func(args ...object.Object) object.Object {
            if len(args) != 0 {
                return newError("wrong number of arguments. want=0. got=%v", len(args))
            }
            return &object.String{Value: "visit https://github.com/ryanlueds/luederlang"}
        },
    },
	"print": &object.Builtin{
		Function: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
