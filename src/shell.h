#include <stdlib.h>
#include <stdio.h>

// Titles
#define SHELL_ERROR_TITLE "ShellError"
#define LEXER_ERROR_TITLE "LexerError"
#define INTERPRETER_ERROR_TITLE "InterpreterError"
#define PARSER_ERROR_TITLE "ParserError"

// Messages
#define ERROR_ARGS_MESSAGE "Invalid arguments..."
#define ERROR_ARGS_MESSAGE_FILENAME "Invalid arguments... To run file must be given a filename"
#define ERROR_FILE_NOT_FOUND "File not found!!!"

// Command Messages
#define HELP_MESSAGE "=--=--=--=--=--=--=[ HELP ]=--=--=--=--=--=--=\n*\n"

// Voids to print messages
#define SHELL_HELP() printf(HELP_MESSAGE);
