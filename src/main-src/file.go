package main

import (
	"errors"
	"io/ioutil"
	"os"
)

func file_exists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func read_file(filename string) string {
	f, err := ioutil.ReadFile(filename)

	if err != nil {
		ece(FILE_ERROR_TITLE, ERROR_FILE_COULDNT_OPEN)
	}

	return string(f)
}
