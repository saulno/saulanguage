package lexer

import (
	"github.com/saulneri1998/saulanguage/token"
)

// Lexer is a tokenizer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a pointer to an instance of Lexer
func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}
	lex.position = lex.readPosition
	lex.readPosition++
}

func (lex *Lexer) peekChar() byte {
	if lex.readPosition >= len(lex.input) {
		return 0
	}
	return lex.input[lex.readPosition]
}

// NextToken returns a Token of the next element in input
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespaces()

	switch lex.ch {
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(lex.ch)}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(lex.ch)}
	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: string(lex.ch)}
	case '/':
		tok = token.Token{Type: token.SLASH, Literal: string(lex.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(lex.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(lex.ch)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(lex.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(lex.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(lex.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(lex.ch)}
	case '=':
		if t, ok := lex.checkBooleanOperator(); ok {
			tok = t
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(lex.ch)}
		}
	case '!':
		if t, ok := lex.checkBooleanOperator(); ok {
			tok = t
		} else {
			tok = token.Token{Type: token.BANG, Literal: string(lex.ch)}
		}
	case '<':
		if t, ok := lex.checkBooleanOperator(); ok {
			tok = t
		} else {
			tok = token.Token{Type: token.LT, Literal: string(lex.ch)}
		}
	case '>':
		if t, ok := lex.checkBooleanOperator(); ok {
			tok = t
		} else {
			tok = token.Token{Type: token.GT, Literal: string(lex.ch)}
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		}
		tok = token.Token{Type: token.ILLEGAL, Literal: string(lex.ch)}
	}

	lex.readChar()
	return tok
}

func (lex *Lexer) readIdentifier() string {
	initial := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[initial:lex.position]
}

func (lex *Lexer) readNumber() string {
	initial := lex.position
	for isDigit(lex.ch) {
		lex.readChar()
	}
	return lex.input[initial:lex.position]
}

func (lex *Lexer) checkBooleanOperator() (token.Token, bool) {
	if lex.peekChar() == '=' {
		ch := lex.ch
		lex.readChar()
		op := string(ch) + string(lex.ch)
		return token.Token{Type: token.LookupBooleanOperator(op), Literal: op}, true
	}
	return token.Token{}, false
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) skipWhitespaces() {
	for lex.ch == ' ' || lex.ch == '\n' || lex.ch == '\t' || lex.ch == '\r' {
		lex.readChar()
	}
}
