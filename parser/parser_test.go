package parser

import (
	"testing"
	"monkey/ast"
	"monkey/lexer"
)

func TestLetStatments(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 33333;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returnd nil")
	}

	if len(program.Statments) != 3 {
		t.Fatalf("Program.Statments does not contain 3 statments. got=%d", len(program.Statments))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statments[i]
		if !testLetStatment(t, stmt, tt.expectedIdentifier ) {
			return
		}
	}
}

func testLetStatment(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' . got=%q", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatment)
	if !ok {
		t.Errorf("s not *ast.LetStatment. got=%T", s)
		return false
	}
	if letStmt.Name.Value !=  name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false

	}

	return true
}

