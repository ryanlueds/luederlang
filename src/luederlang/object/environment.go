package object

func NewEnvironment() *Environment {
    s := make(map[string]Object)
    return &Environment{store: s}
}

type Environment struct {
    store map[string]Object;
}


func (e *Environment) Set(name string, value Object) Object {
    e.store[name] = value 
    return value
}

func (e *Environment) Get(name string) (Object, bool) {
    val, ok := e.store[name]
    return val, ok
}
