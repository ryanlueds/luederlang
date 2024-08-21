package token

import (
    "strings"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"                 // add, foobar, x, y, ...
    INT_LITERAL   = "INT_LITERAL"   // 1343456
    FLOAT_LITERAL = "FLOAT_LITERAL"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
    INT      = "INT"
    FLOAT    = "FLOAT" 
    BOOL     = "BOOL"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
    "int":    INT,
    "float":  FLOAT,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func LookupNumber(number string) TokenType {
    count := strings.Count(number, ".")
    var tok TokenType
    switch count {
    case 0:
        tok = INT_LITERAL
    case 1:
        tok = FLOAT_LITERAL
    default:
        tok = ILLEGAL
    }
    return tok
}
