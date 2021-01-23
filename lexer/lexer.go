package lexer

import (
	"github.com/subintp/monkey/token"
)

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

// NextToken - get next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

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
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok = token.New(string(token.LookupIdent(literal)), literal)
			return tok
		} else if isDigit(l.ch) {
			tok = token.New(token.INT, l.readNumber())
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
			return tok
		}
	}
	l.readChar()
	return tok
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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
