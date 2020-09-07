package token

// Constant types of tokens
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	ASSIGN    = "="
	BANG      = "!"
	LT        = "<"
	GT        = ">"
	LEQ       = "<="
	GEQ       = ">="
	EQ        = "=="
	NOTEQ     = "!="
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

var booleanOperators = map[string]Type{
	"==": EQ,
	"!=": NOTEQ,
	"<":  LT,
	">":  GT,
	"<=": LEQ,
	">=": GEQ,
}

// LookupIdentifier checks if argument is a keyword or a user defined identifier
func LookupIdentifier(ident string) Type {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}

// LookupBooleanOperator searches which boolean operator corresponds to argument
func LookupBooleanOperator(op string) Type {
	return booleanOperators[op]
}

// Type is a descriptor of the kind of tokens
type Type string

// Token is th ebasic element of the lexer
type Token struct {
	Type    Type
	Literal string
}
