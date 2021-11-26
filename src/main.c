#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "error.h"
#include "shell.h"

int main(int argc, char **argv) {
    if (argc >= 2) {
        if (strcmp(argv[1], "--help") == 0 || strcmp(argv[1], "help") == 0 || strcmp(argv[1], "--h") == 0) {
            SHELL_HELP();
        } else if (strcmp(argv[1], "run")) {
            if (argc >= 3) {
                
            }
        }
    } else {
        struct Error error_args;

        error_args.title = SHELL_ERROR_TITLE;
        error_args.message = ERROR_ARGS_MESSAGE;

        error_debug(error_args);
    }   
}