package token

type tokenType string

type Token struct {
	Type 	tokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"
	
	IDNET = "IDENT"
	INT	  = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)