#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void interpreter_run(char* filename) {
    struct filereadreturn val = read_file(filename);
    
    lexer(val.line, val.linesCount);
}
