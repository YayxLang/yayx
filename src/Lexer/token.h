#include <stdio.h>

struct Token {
    char* Type;
    char* Value;
    int LineNumber;
    int ColumnNumber;
} Token;

struct Token newtoken(char* Type, char* Value, int LineNumber, int ColumnNumber) {
    struct Token tk;

    tk.Type = Type;
    tk.Value = Value;
    tk.LineNumber = LineNumber;
    tk.ColumnNumber = ColumnNumber;

    return tk;
}

// Operators
#define PLUS "+"
#define MINUS "-"
#define MUL "*"
#define DIV "/"
#define BANG "!"

// Logical Operators
#define GT ">"
#define GTE ">="
#define LT "<"
#define LTE "<="
#define ASN "="
#define EQ "=="
#define NEQ "!="

// Btiwise operators
#define OROR "||"
#define ANDAND "&&"

// Structural tokens
#define OPENPARENT "("
#define CLOSEPARENT ")"
#define OPENBRACKET "["
#define CLOSEBRACKET "]"
#define OPENBRACE "{"
#define CLOSEBRACE "}"

// Seperators
#define COMMA ","
#define DOT "."
#define SEMICOL ";"

// Types
#define IDENTIFIER "Identifier"
#define STRING "String"
#define INT "Integer"
