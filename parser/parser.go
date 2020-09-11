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
	return nil
}
