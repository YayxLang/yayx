package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	valueFile := string(f)
	splitedValue := strings.Split(valueFile, "\n")

	for i := 0; i < len(splitedValue); i++ {
		fmt.Printf(splitedValue[i])
	}


	return string(f)
}
