package main

import (
	"fmt"
    "io"
    "os"
	"luederlang/repl"
	"luederlang/lexer"
	"luederlang/parser"
    "luederlang/evaluator"
    "luederlang/object"
)

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func executeFile(file string) {
    env := object.NewEnvironment()
    l := lexer.New(file)
    p := parser.New(l)

    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
        for _, msg := range p.Errors() {
            io.WriteString(os.Stdout, "\t"+msg+"\n")
        }
        return
    }

    eval := evaluator.Eval(program, env)
    // TODO: Delete this
    if eval != nil {
        io.WriteString(os.Stdout, eval.Inspect())
        io.WriteString(os.Stdout, "\n")
    }
}

func main() {
    args := os.Args[1:]
    switch len(args) {
    case 0:
        fmt.Printf("type help() for help\n")
        repl.Start(os.Stdin, os.Stdout)
    case 1:
        bytes, err := os.ReadFile(args[0])
        if err != nil {
            fmt.Println(err)
        }
        executeFile(string(bytes))
    default:
        fmt.Println(fmt.Errorf("Too many arguments: %v\n", args))
    } 
}
