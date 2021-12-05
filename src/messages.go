package main

import "fmt"

const (
	VERSION = "1.0.0 BETA"
)

const (
	// Titles of errors
	SHELL_ERROR_TITLE       = "ShellError"
	FILE_ERROR_TITLE        = "FileError"
	LEXER_ERROR_TITLE       = "LexerError"
	INTERPRETER_ERROR_TITLE = "InterpreterError"
	PARSER_ERROR_TITLE      = "ParserError"

	// Message of errors
	ERROR_ARGS_MESSAGE           = "Invalid arguments..."
	ERROR_ARGS_MESSAGE_FILENAME  = "Invalid arguments... To run file must be given a filename"
	ERROR_FILE_NOT_FOUND         = "File not found!!!"
	ERROR_FILE_COULDNT_OPEN      = "File could not be open!!!"
	ERROR_LEXER_FLOAT_DOT        = "Float has already '.' inside of it!"
	ERROR_LEXER_PARENTS_OPENDED  = "Parents ('(') opened but not closed"
	ERROR_LEXER_BRACE_OPENDED    = "Brace ('{') opened but not closed"
	ERROR_LEXER_BRACKETS_OPENDED = "Brackets ('[') opened but not closed"
	ERROR_LEXER_UNKNOWN_CHAR     = "Unknow Character ("

	// Command Messages
	HELP_MESSAGE    = "=--=--=--=--=--=--=[ HELP ]=--=--=--=--=--=--=\n*\n"
	VERSION_MESSAGE = "Version of Yayx interpreter: " + colorYellow + VERSION + colorReset
)

func shell_help()    { fmt.Println(HELP_MESSAGE) }
func shell_version() { fmt.Println(VERSION_MESSAGE) }
