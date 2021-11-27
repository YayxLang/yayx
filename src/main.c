#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "error.h"
#include "shell.h"
#include "file.h"
#include "console.h"
#include "interpreter.h"

// Main entry point of the application
int main(int argc, char **argv) {
    if (argc >= 2) {
        if (strcmp(argv[1], "--help") == 0 || strcmp(argv[1], "help") == 0 || strcmp(argv[1], "--h") == 0) {
            SHELL_HELP();
        } else if (strcmp(argv[1], "run") == 0) {
            if (argc >= 3) {
                int fileex = file_exists(argv[2]);

                if (fileex == 1) {
                    interpreter_run(argv[2]);
                } else {
                    struct Error error_args_filename = ece(ERROR_FILE_NOT_FOUND, SHELL_ERROR_TITLE);
                }
            } else {
                struct Error error_args_run = ece(ERROR_ARGS_MESSAGE_FILENAME, SHELL_ERROR_TITLE);
            }
        }
    } else {
        struct Error error_args = ece(ERROR_ARGS_MESSAGE, SHELL_ERROR_TITLE);
    }   
}