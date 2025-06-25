package lexer_test

import (
    "io/ioutil"
    "testing"
    "poncho/token"
    "poncho/lexer"
)

func TestLexerFromFile(t *testing.T) {
    inputBytes, err := ioutil.ReadFile("../../scripts/poncho.p")

    if err != nil {
        t.Fatalf("failed to read input file: %v", err)
    }

    input := string(inputBytes)
    lexer := lexer.New(input)

    expectedTokens := []token.Token{
        {Type: token.CONST, Literal: "const"},
        {Type: token.IDENT, Literal: "x"},
        {Type: token.ASSIGN, Literal: "="},
        {Type: token.INT, Literal: "10"},
        {Type: token.LOG, Literal: "log"},
        {Type: token.LPAREN, Literal: "("},
        {Type: token.IDENT, Literal: "x"},
        {Type: token.RPAREN, Literal: ")"},
        {Type: token.EOF, Literal: ""},
    }

    for i, expected := range expectedTokens {
        tok := lexer.NextToken()

        if tok.Type != expected.Type || tok.Literal != expected.Literal {
            t.Fatalf("token %d wrong. expected={Type:%q Literal:%q}, got={Type:%q Literal:%q}", i, expected.Type, expected.Literal, tok.Type, tok.Literal)
        }
    }
}


