package parser

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	NUMBER
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
)

type Token struct {
	Type    TokenType
	Literal string
}
