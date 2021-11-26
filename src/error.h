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