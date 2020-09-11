package ast

import "github.com/saulneri1998/saulanguage/token"

// Node is the basic element of our AST
type Node interface {
	TokenLiteral() string
}

// Statement is a Node of our AST, all statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// Expression is a Node of our AST, all expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

// Program is the root Node for every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral is a method to implement Node interface, used for debugging and testing
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return " "
}

// Statements

// LetStatement id a node that reoresents a variable declaration
type LetStatement struct {
	Token token.Token // the token.LET token Value string
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral is a method to implement Node interface, used for debugging and testing
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier is a Node to represent the name of a variable
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) statementNode() {}

// TokenLiteral is a method to implement Node interface, used for debugging and testing
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
