package lexer

import (
	"testing"

	"github.com/eduardolat/edulang/token"
)

func TestNextToken_Basic(t *testing.T) {
	input := "+(){};,="

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.ASSIGN, "="},
	}

	lx := New(input)

	for i, tt := range tests {
		tok := lx.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. Expected: %q, got: %q",
				i, tt.expectedType, tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. Expected: %q, got: %q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}
}
