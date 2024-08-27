# Luederlang v0.0.1
## About:
Luederlang is an interpreted language written in Go. Largely inspired
by [Thorsten Ball's implementation](https://interpreterbook.com/) with key changes.
Incredibly minimal to fit the hard deadline of the school semester starting. I 
like naming things after myself, and Luederlang was too good of a name to pass up.
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
// A little verbose: we dont have loops but we have recursion
let fizzbuzz = fun(i, n) {
    if (i > n) {
        print("\n");
        return 0;
    }

    // A little verbose: we don't have else-ifs
    if (i % 3 == 0) { print("fizz"); }
    if (i % 5 == 0) { print("buzz"); }
    if (i % 3 != 0 && i % 5 != 0) { print(i); }
    print(" ");

    fizzbuzz(i + 1, n);
}

fizzbuzz(1, 30);
```
```
~/ go run main.go ryan.lueder
1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19 buzz fizz 22 23 fizz buzz 26 fizz 28 29 fizzbuzz
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
