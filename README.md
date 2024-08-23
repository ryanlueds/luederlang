# Luederlang v0.0.1
## About:
Luederlang is an interpreted language written in Go. Largely inspired
by [Thorsten Ball's implementation](https://interpreterbook.com/) with key changes.
Incredibly minimal to fit the hard deadline of the school semester starting.
## Specs
- Lexer (A5 Military Grade Wagyu)
- Parser (Top Down Operator Precedence / Pratt Parser)
- Evaluator (Treewalking Evaluator)
## Features
- C like sytax
- Integers, Booleans, Floats, String literals
- Comments
- Upcasting infix expressions based on operator
- First class and higher-order functions
- Builtin Functions (print, len, help)
- REPL
## Missing Features
- Arrays
- Maps
- Loops
## Examples
### Fizzbuzz without loops and without else-if's:
```
// ryan.lueder
int number = 15

if(number % 3 == 0 && number % 5 == 0) {
    print("fizzbuzz")
} else {
    if(number % 3 == 0) {
        print("fizz")
    }
    if(number % 5 == 0) {
        print("buzz")
    }
}
```
```
~/ go run main.go ryan.lueder
fizzbuzz
```
Note that `ryan.lueder` is the entry point for Luederlang files. This is not enforced by the interpreter but IS industry standard and IS largely believed to be best code style.

### Fixed modulus division:
`-3 % 5 // => 2, not -3 like in C or Golang`
### Higher order functions:
```
// ryan.lueder
// This is the "Hello, World!" program equivalent in Luederlang
let messageUser = fun(msg) {
    return fun(user) {
        msg + " " + user // return not required
    }
}

let helloUser = messageUser("Hello")
print(helloUser("Ryan"))
```
```
~/ go run main.go ryan.lueder
Hello Ryan
```
### Typing:
```
float x = 7.8 // ok
let x = 7.8   // ok
```
### REPL:
```
~/ go run main.go
type help() for help
>> 5 + 55.5
60.5
>>
```
