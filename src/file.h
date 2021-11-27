#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define M1 2048
#define M2 2048

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

// Read file contents and return it
char** read_file(char* filename) {
    FILE* file = fopen(filename, "r");
    char line[M2]*;
    int k = 0;
    int no = 0;

    if (file == NULL) {
        struct Error error_args_run = ece(ERROR_FILE_COULDNT_OPEN, FILE_ERROR_TITLE);
        exit(0);
    } else {
        while (fgets(line[k], M1, file)) {
            line[k][strlen(line[k])-1] = '\0';
            k++;
        }

        fclose(file);
    }

    return line;
}

// Write into file content and return it value
char* write_file(char* filename, char* data) {

}
