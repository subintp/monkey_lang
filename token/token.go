package token

// Type is string
type Type string

// Token in lexer
type Token struct {
	Type    Type
	Literal string
}

// Possible Tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	ADD    = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// New - create new tocken
func New(tokenType string, literal string) Token {
	return Token{
		Type:    Type(tokenType),
		Literal: literal,
	}
}

// LookupIdent = check if identifier or keyword
func LookupIdent(ident string) Type {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}
