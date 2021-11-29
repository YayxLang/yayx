package main

type Token struct {
	Type       string // Type of token
	Value      string // Value of the token
	LineColumn int    // Line of token
}

func createToken(Type string, Value string, LineColumn int) Token {
	var t Token

	t.Type = Type
	t.Value = Value
	t.LineColumn = LineColumn

	return t
}

const (
	// Operators
	PLUS  = "+"
	MINUS = "-"
	MUL   = "*"
	DIV   = "/"
	BANG  = "!"

	// Logical Operators
	GT  = ">"
	GTE = ">="
	LT  = "<"
	LTE = "<="
	ASN = "="
	EQ  = "=="
	NEQ = "!="

	// Btiwise operators
	OROR   = "||"
	ANDAND = "&&"

	// Structural tokens
	OPENPARENT   = "("
	CLOSEPARENT  = ")"
	OPENBRACKET  = "["
	CLOSEBRACKET = "]"
	OPENBRACE    = "{"
	CLOSEBRACE   = "}"

	// Seperators
	COMMA   = ","
	DOT     = "."
	SEMICOL = ";"

	// Types
	IDENTIFIER = "Identifier"
	STRING     = "String"
	INT        = "Integer"
	FLOAT      = "Float64"
)

const (
	// Operators
	TOKENNAME_PLUS  = "Plus"
	TOKENNAME_MINUS = "Minus"
	TOKENNAME_MUL   = "Multiply"
	TOKENNAME_DIV   = "Divide"
	TOKENNAME_BANG  = "Bang"

	// Logical Operators
	TOKENNAME_GT  = "GT"
	TOKENNAME_GTE = "GTE"
	TOKENNAME_LT  = "LT"
	TOKENNAME_LTE = "LTE"
	TOKENNAME_ASN = "ASN"
	TOKENNAME_EQ  = "EQ"
	TOKENNAME_NEQ = "NEQ"

	// Btiwise operators
	TOKENNAME_OROR   = "OROR"
	TOKENNAME_ANDAND = "ANDAND"

	// Structural tokens
	TOKENNAME_OPENPARENT   = "OPENPARENT"
	TOKENNAME_CLOSEPARENT  = "CLOSEPARENT"
	TOKENNAME_OPENBRACKET  = "OPENBRACKET"
	TOKENNAME_CLOSEBRACKET = "CLOSEBRACKET"
	TOKENNAME_OPENBRACE    = "OPENBRACE"
	TOKENNAME_CLOSEBRACE   = "CLOSEBRACE"

	// Seperators
	TOKENNAME_COMMA   = "COMMA"
	TOKENNAME_DOT     = "DOT"
	TOKENNAME_SEMICOL = "SEMICOL"
)

const (
	CONSTANTS_LETTER = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	CONSTANTS_NUMBER = "0123456789"
)
