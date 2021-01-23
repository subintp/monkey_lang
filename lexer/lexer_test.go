package lexer

import (
	"testing"

	"github.com/subintp/monkey/token"
)

func TestNextToken(t *testing.T) {
	code := "=+(){},;"

	expectedTokens := []token.Token{
		{token.ASSIGN, "="},
		{token.ADD, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(code)

	for i, expectedToken := range expectedTokens {
		token := l.NextToken()

		if token.Type != expectedToken.Type {
			t.Fatalf("test[%d] - wrong token type. expected=%q got=%q", i, expectedToken.Type, token.Type)
		}

		if token.Literal != expectedToken.Literal {
			t.Fatalf("test[%d] - wrong literal . expected=%q got=%q", i, expectedToken.Literal, token.Literal)
		}
	}
}
