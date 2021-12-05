package main

import (
	"fmt"
	"os"
)

// Structure thats create a error object
type Error struct {
	message string // Message of error
	title   string // Title of error
}

// Void that will print out a error structure
func error_debug(error *Error) {
	fmt.Println("[" + colorRed + error.title + colorReset + "]: " + error.message)
	os.Exit(0)
}

func ece(message string, title string) *Error {
	var err = new(Error)

	err.message = message
	err.title = title

	error_debug(err)

	return err
}
