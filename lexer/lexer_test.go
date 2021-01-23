package lexer

import (
	"testing"

	"github.com/subintp/monkey/token"
)

func TestNextToken(t *testing.T) {
	code := `
		let five = 5;
		let ten = 10;
		let add = fn(x,y) {
			x+y;
		};
		let result = add(five, ten);
	`

	expectedTokens := []token.Token{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
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
