package main

import (
    "fmt"
    "poncho/lexer"
)

func main() {
    input := `const x = 10
log(x)          # prints 10`

    l := lexer.New(input)

    for {
        tok := l.NextToken()
        fmt.Printf("%+v\n", tok)
        if tok.Type == "EOF" {
            break
        }
    }
}

