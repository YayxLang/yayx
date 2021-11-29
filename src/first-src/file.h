#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define M1 130
#define M2 20

// Is file exists?
// Returns 0 or 1 (0 - Not exists | 1 - File exists)
int file_exists(char* filename) {
    int return_state = 0;

    FILE* file = fopen(filename, "r");
    
    if (file != NULL) {
        return_state = 1;
    } else {
        return_state = 0;
    }

    return return_state;
}

struct filereadreturn {
    char line[M2][M1];
    int linesCount;
} filereadreturn;

// Read file contents and return it
void read_file(char* filename) {
    FILE* file = fopen(filename, "r");
    char line[M2][M1];
    int k = 0;
    int no = 0;

    if (file == NULL) {
        ece(ERROR_FILE_COULDNT_OPEN, FILE_ERROR_TITLE);
        exit(0);
    }

    while (fgets(line[k], M1, file)) {
        line[k][strlen(line[k])-1] = '\0';
        k++;
    }
    

    fclose(file);

    // struct filereadreturn returnstuctor;
    // returnstuctor.line = line;
    // returnstuctor.linesCount = k;

    // return returnstuctor;
}

// Write into file content and return it value
char* write_file(char* filename, char* data) {

}
