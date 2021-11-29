#include <stdio.h>
#include <stdlib.h>

// This function is used from stackoverflow heh thanks a lot
// https://stackoverflow.com/questions/12939370/c-appending-char-to-char
int append(char*s, size_t size, char c) {
    if(strlen(s) + 1 >= size) {
        return 1;
    }
    int len = strlen(s);
    s[len] = c;
    s[len+1] = '\0';
    return 0;
}

char* appendchar(char* str, char* toappend) {
    char* toreturn = str;

    for (int i = 0; i < strlen(str); i++) {
        append(toreturn, strlen(toreturn),toappend[i]);
    }

    printf("%s", toreturn);
}

int arraylen(char** array, size_t len) {
    int i = 0;

    for (size_t i = 0; i < len; i++) {

    }

    return i;
}
