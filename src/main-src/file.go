package main

import (
	"bufio"
	"errors"
	"os"
)

func file_exists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

// ====================================================================
//
// Old function of reading a file:
//
// ====================================================================
//func readFile(filename string) string {
//	f, err := ioutil.ReadFile(filename)
//
//	if err != nil {
//		ece(FILE_ERROR_TITLE, ERROR_FILE_COULDNT_OPEN)
//	}
//
//	valueFile := string(f)
//	splitedValue := strings.Split(valueFile, "\n")
//
//	for i := 0; i < len(splitedValue); i++ {
//		fmt.Printf(splitedValue[i])
//	}
//
//
//	return string(f)
//}
// ====================================================================

// func valueFileSplit(valueFile string) string {
// 	var splited []string = strings.Split(valueFile, "\n")
// 	var endVal string = ""

// 	for i := 0; i < len(splited); i++ {
// 		endVal = endVal + strings.ReplaceAll(splited[i], "\n", "") + " "
// 	}

// 	// ====================================================================
// 	// Debug function
// 	// ====================================================================
// 	//for i := 0; i < len(splited); i++ {
// 	//fmt.Printf(splited[i])
// 	//}
// 	// ====================================================================

// 	return endVal
// }

var (
	fileVal string
)

func readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		ece(FILE_ERROR_TITLE, ERROR_FILE_COULDNT_OPEN)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	if err := scanner.Err(); err != nil {
		ece(FILE_ERROR_TITLE, ERROR_FILE_COULDNT_OPEN)
	}

	for scanner.Scan() {
		// fileVal := valueFileSplit(scanner.Text())

		fileVal += scanner.Text()
	}
	interpreter_run(fileVal)
}
