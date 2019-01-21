package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser{
	p := &Parser{
		l: l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statments = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatment()
		if stmt != nil {
			program.Statments = append(program.Statments, stmt)
		}
		p.nextToken()
	}
	return program
}


func (p *Parser) parseStatment() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatment()
	case token.RETURN:
		return p.parseReturnStatment()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatment() *ast.LetStatment {
	stmt := &ast.LetStatment{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s insted", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatment() *ast.ReturnStatment {
	stmt := &ast.ReturnStatment{Token: p.curToken}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
