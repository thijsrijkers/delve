package token

import (
    "strings"
)

// TokenType is a string alias for identifying token categories
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
    // Special tokens
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers + literals
    IDENT = "IDENT"
    INT   = "INT"

    // Operators
    ASSIGN = "="

    // Delimiters
    LPAREN = "("
    RPAREN = ")"

    // Keywords
    CONST = "CONST"
    LOG   = "LOG"
)

var keywords = map[string]TokenType{
	"const": CONST,
	"log":   LOG,
}

func LookupIdent(ident string) TokenType {
	identLower := strings.ToLower(ident) // normalize to lowercase
	if tok, ok := keywords[identLower]; ok {
		return tok
	}
	return IDENT
}
