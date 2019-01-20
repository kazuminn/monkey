package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statmentNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statments []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatment struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatment) statmentNode() {}
func (ls *LetStatment) TokenLiteral() string { return ls.Token.Literal}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal}


type ReturnStatment struct {
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statmentNode()	()
func (rs *ReturnStatment) TokenLiteral() string { return rs.Token.Literal}
