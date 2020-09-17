package parser

import (
	"github.com/saulneri1998/saulanguage/ast"
	"github.com/saulneri1998/saulanguage/lexer"
	"github.com/saulneri1998/saulanguage/token"
)

// Parser constructs an AST out of a lexer input
type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New creates and initializes a new Parser
func New(lex *lexer.Lexer) Parser {
	p := Parser{lex: lex}

	// Reading two tokens, initializes our both tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

// ParseProgram constructs the AST
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	if !p.expectPeek(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}
