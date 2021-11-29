//  Includes
#include <stdio.h>
#include <stdlib.h>

// Structure thats create a error object
struct Error {
    char* message; // Message of error
    char* title; // Title of error
} Error;

// Void that will print out a error structure
void error_debug(struct Error error) {
    // Print out
    printf("[%s]: %s\n", error.title, error.message);
}

// Less lines of code and easier creating errors
struct Error ece(char* message, char* title) {
    struct Error ece_error;

    ece_error.title = title;
    ece_error.message = message;

    error_debug(ece_error);
    
    return ece_error;
}

