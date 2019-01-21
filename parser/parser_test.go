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
	checkParserErrors(t,p)
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatments(t *testing.T) {
	input := `
	return 5;
	return 20;
	return 999999;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statments) != 3 {
		t.Fatalf("program.Statments does not contain 3 statments, got=%d", len(program.Statments))
	}

	for _, stmt := range program.Statments {
		returnStmt, ok := stmt.(*ast.ReturnStatment)
		if !ok {
			t.Errorf("stmt not *ast.returnStatment. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",returnStmt.TokenLiteral())
		}
	}
}
