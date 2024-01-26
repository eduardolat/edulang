package token

// TokenType is a string that represents the type of a token.
type TokenType string

// Token is a struct that holds the type of a token and its literal value.
type Token struct {
	Type    TokenType
	Literal string
}

// New returns a new Token.
func New(tokenType TokenType, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

// NewFromByte returns a new Token from a byte literal.
func NewFromByte(tokenType TokenType, literal byte) Token {
	return Token{Type: tokenType, Literal: string(literal)}
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL" // A token that we don't know about
	EOF     = "EOF"     // The end of the file

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1234567890
	STRING = "STRING" // "foobar"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent returns the TokenType for a given identifier.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
