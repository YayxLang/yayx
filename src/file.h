#include <stdio.h>
#include <string.h>
#include <stdlib.h>

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
char* read_file(char* filename) {

}

// Write into file content and return it value
char* write_file(char* filename, char* data) {

}