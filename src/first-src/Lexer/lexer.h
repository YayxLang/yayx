#include <stdlib.h>
#include <stdio.h>
#include <string.h>

void lexer(char valuefile[M2][M1], int linesCount) {
    char currentCharacter;
    struct Token* tokens;
    int space = 0;
    int stringOpened = 0;
    int no = 0;
    no = linesCount;
    for (linesCount = 0; linesCount < no; linesCount++) {
        printf("%s", valuefile[linesCount]);
    }
}