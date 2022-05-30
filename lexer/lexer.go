package lexer

import (
	"github.com/kubabialy/donkey/token"
)

type Lexer struct {
    input        string
    position     int  // current position in input (point to current char)
    readPosition int  // current reading position in input (after current char)
    ch           byte // current char under examination
}

func New(input string) *Lexer {
    lxr := &Lexer{input: input}
    lxr.readChar()
    return lxr
}

func (lxr *Lexer) NextToken() token.Token {
    var tok token.Token

    lxr.skipWhitespace()

    switch lxr.ch {
    case '!':
        if lxr.peekChar() == '=' {
            ch := lxr.ch
            lxr.readChar()
            literal := string(ch) + string(lxr.ch)
            tok = token.Token{Type: token.NOT_EQ, Literal: literal}
        } else {
            tok = newToken(token.BANG, lxr.ch)
        }
    case '/':
        tok = newToken(token.SLASH, lxr.ch)
    case '*':
        tok = newToken(token.ASTERISK, lxr.ch)
    case '>':
        tok = newToken(token.GT, lxr.ch)
    case '<':
        tok = newToken(token.LT, lxr.ch)
    case '=':
        if lxr.peekChar() == '=' {
            ch := lxr.ch
            lxr.readChar()
            literal := string(ch) + string(lxr.ch)
            tok = token.Token{Type: token.EQ, Literal: literal}
        } else {
            tok = newToken(token.ASSIGN, lxr.ch)
        }
    case ';':
        tok = newToken(token.SEMICOLON, lxr.ch)
    case '(':
        tok = newToken(token.LPAREN, lxr.ch)
    case ')':
        tok = newToken(token.RPAREN, lxr.ch)
    case ',':
        tok = newToken(token.COMMA, lxr.ch)
    case '+':
        tok = newToken(token.PLUS, lxr.ch)
    case '-':
        tok = newToken(token.MINUS, lxr.ch)
    case '{':
        tok = newToken(token.LBRACE, lxr.ch)
    case '}':
        tok = newToken(token.RBRACE, lxr.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(lxr.ch) {
            tok.Literal = lxr.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        } else if isDigit(lxr.ch) {
            tok.Type = token.INT
            tok.Literal = lxr.readNumber()
            return tok
        } else {
            tok = newToken(token.ILLEGAL, lxr.ch)
        }
    }

    lxr.readChar()
    return tok
}

func (lxr *Lexer) readIdentifier() string {
    position := lxr.position
    for isLetter(lxr.ch) {
        lxr.readChar()
    }
    return lxr.input[position:lxr.position]
}

// verifies if provided character is considered by the language to be a letter
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

// create new token based on provided char and its type
func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
 `position` and `readPosition` are used to access characters in input by using them
 as index, e.g `lxr.input[lxr.readPosition]`. The reason for these two "pointers"
 pointing into our input srting is the fact that we will need to be able to "peek"
 further into the input and look after the current character to see what comes up next.
*/
func (lxr *Lexer) readChar() {
    if lxr.readPosition >= len(lxr.input) {
        lxr.ch = 0
    } else {
        lxr.ch = lxr.input[lxr.readPosition]
    }

    lxr.position = lxr.readPosition
    lxr.readPosition += 1
}

/*
`peekChar()` is really similar to `readChar()` , except that it doesn’t increment l.position and
l.readPosition . We only want to “peek” ahead in the input and not move around in it, so
we know what a call to `readChar()` would return.
*/
func (lxr *Lexer) peekChar() byte {
    if (lxr.readPosition >= len(lxr.input)) {
        return 0
    } else {
        return lxr.input[lxr.readPosition]
    }
}

func (lxr *Lexer) readNumber() string {
    position := lxr.position
    for isDigit(lxr.ch) {
        lxr.readChar()
    }

    return lxr.input[position:lxr.position]
}

/*
because whitespaces are just noise not required to parse the code
`Lexer` will catch whitespace/new lines and move the pointer forward
*/
func (lxr *Lexer) skipWhitespace() {
    for lxr.ch == ' ' || lxr.ch == '\t' || lxr.ch == '\n' || lxr.ch == '\r' {
        lxr.readChar()
    }
}

