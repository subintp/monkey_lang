package lexer

import "github.com/subintp/monkey/token"

// Lexer - represents lexer internal state
type Lexer struct {
	input        string
	position     int // current input position
	readPosition int // current reading position
	ch           byte
}

// New - create new lexer with input
func New(input string) *Lexer {
	lexer := Lexer{input: input}
	lexer.readChar()
	return &lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken = get next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, string(l.ch))
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch))
	case '(':
		tok = token.New(token.LPAREN, string(l.ch))
	case ')':
		tok = token.New(token.RPAREN, string(l.ch))
	case ',':
		tok = token.New(token.COMMA, string(l.ch))
	case '+':
		tok = token.New(token.ADD, string(l.ch))
	case '{':
		tok = token.New(token.LBRACE, string(l.ch))
	case '}':
		tok = token.New(token.RBRACE, string(l.ch))
	case 0:
		tok = token.New(token.EOF, "")
	}
	l.readChar()
	return tok
}
