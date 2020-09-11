package parser

import (
	"testing"

	"github.com/saulneri1998/saulanguage/ast"
	"github.com/saulneri1998/saulanguage/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
let y = 10
let foo = 12345`

	var lex *lexer.Lexer = lexer.New(input)
	var p Parser = New(lex)

	var program *ast.Program = p.ParseProgram()
	if program == nil {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, expected string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != expected {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", expected, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != expected {
		t.Errorf("s.Name not '%s'. got=%s", expected, letStmt.Name)
		return false
	}
	return true
}
