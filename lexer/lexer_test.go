package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	test := []struct{
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS,   "+"},
		{token.RPAREN, "("},
		{token.LPAREN, ")"},


	}
}