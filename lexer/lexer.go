package lexer

import (
    "poncho/token"
    "unicode"
)

// Lexer processes the input string and produces tokens sequentially.
// It keeps track of its current reading position and the character being examined.
type Lexer struct {
    input        string // source code to tokenize
    currentPos   int    // current index in input, points to currentChar
    nextReadPos  int    // index to read next char from input (lookahead)
    currentChar  byte   // current char under examination
}

// New creates and initializes a Lexer to start tokenizing input from the beginning.
func New(input string) *Lexer {
    lexer := &Lexer{input: input}
    lexer.readChar() // initialize currentChar to first character
    return lexer
}

// readChar advances the lexer's positions by one and updates currentChar.
// Sets currentChar to 0 to signal EOF when input ends.
// This allows easy lookahead without losing track of current char.
func (lexer *Lexer) readChar() {
    if lexer.nextReadPos >= len(lexer.input) {
        lexer.currentChar = 0 // EOF sentinel
    } else {
        lexer.currentChar = lexer.input[lexer.nextReadPos]
    }
    lexer.currentPos = lexer.nextReadPos
    lexer.nextReadPos++
}

// newToken creates a token from a single character lexeme,
// used for simple one-character tokens for clarity and reuse.
func newToken(tokenType token.TokenType, char byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(char)}
}

// skipWhitespace ignores spaces, tabs, and newlines,
// since they don’t carry meaning in token stream.
func (lexer *Lexer) skipWhitespace() {
    for lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\n' || lexer.currentChar == '\r' {
        lexer.readChar()
    }
}

// skipComment ignores all characters until end of line or EOF,
// since comments do not affect the semantics of the code.
func (lexer *Lexer) skipComment() {
    for lexer.currentChar != '\n' && lexer.currentChar != 0 {
        lexer.readChar()
    }
}

// isLetter identifies characters that can start or be part of identifiers,
// allowing letters and underscore for idiomatic naming.
func isLetter(character byte) bool {
    return unicode.IsLetter(rune(character)) || character == '_'
}

// isDigit identifies decimal digit characters,
// used to recognize integer literals.
func isDigit(character byte) bool {
    return '0' <= character && character <= '9'
}

// readIdentifier reads an entire identifier or keyword token,
// preventing partial tokenization of multi-character names.
func (lexer *Lexer) readIdentifier() string {
    startPosition := lexer.currentPos
    for isLetter(lexer.currentChar) || isDigit(lexer.currentChar) {
        lexer.readChar()
    }
    return lexer.input[startPosition:lexer.currentPos]
}

// readNumber reads a full multi-digit integer token,
// so numbers like '1234' become one token, not four.
func (lexer *Lexer) readNumber() string {
    startPosition := lexer.currentPos
    for isDigit(lexer.currentChar) {
        lexer.readChar()
    }
    return lexer.input[startPosition:lexer.currentPos]
}


// NextToken returns the next valid token from the input stream,
// skipping whitespace and comments to focus on meaningful tokens.
// It handles single-char tokens immediately for simplicity,
// and delegates reading of identifiers and numbers to helper methods.
func (lexer *Lexer) NextToken() token.Token {
    lexer.skipWhitespace()

    // Skip comments entirely since they don't affect parsing
    if lexer.currentChar == '#' {
        lexer.skipComment()
        lexer.skipWhitespace()
    }

    var nextToken token.Token

    // Recognize single-character tokens right away
    switch lexer.currentChar {
    case '=':
        nextToken = newToken(token.ASSIGN, lexer.currentChar)
    case '(':
        nextToken = newToken(token.LPAREN, lexer.currentChar)
    case ')':
        nextToken = newToken(token.RPAREN, lexer.currentChar)
    case 0:
        nextToken = token.Token{Type: token.EOF, Literal: ""} // signal end of input
    default:
        // If letter, read full identifier or keyword to avoid splitting multi-char names
        if isLetter(lexer.currentChar) {
            literal := lexer.readIdentifier()
            nextToken.Type = token.LookupIdent(literal) // distinguish keywords from user names
            nextToken.Literal = literal
            return nextToken
        } else if isDigit(lexer.currentChar) {
            // Numbers can have multiple digits, read them all as one token
            literal := lexer.readNumber()
            nextToken.Type = token.INT
            nextToken.Literal = literal
            return nextToken
        } else {
            // Unknown characters are illegal; lexer marks them for parser to handle errors
            nextToken = newToken(token.ILLEGAL, lexer.currentChar)
        }
    }

    lexer.readChar() // move forward after single-char tokens or illegal ones
    return nextToken
}

