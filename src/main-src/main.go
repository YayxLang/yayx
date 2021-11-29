package main

import (
	"os"
)

func main() {
	argv := os.Args

	if len(argv) >= 2 {
		if argv[1] == "--help" || argv[1] == "help" || argv[1] == "--h" || argv[1] == "-help" || argv[1] == "-h" {
			shell_help()
		} else if argv[1] == "run" {
			if len(argv) >= 3 {
				fileExists := file_exists(argv[2])

				if fileExists == true {
					interpreter_run(argv[2])
				} else {
					ece(ERROR_FILE_NOT_FOUND, SHELL_ERROR_TITLE)
				}
			} else {
				ece(ERROR_ARGS_MESSAGE_FILENAME, SHELL_ERROR_TITLE)
			}
		} else if argv[1] == "version" || argv[1] == "--version" || argv[1] == "--v" || argv[1] == "-version" || argv[1] == "-v" {
			shell_version()
		}
	} else {
		ece(ERROR_ARGS_MESSAGE, SHELL_ERROR_TITLE)
	}
}
