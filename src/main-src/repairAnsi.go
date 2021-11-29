package main

import (
	"os"

	"golang.org/x/sys/windows"
)

// This function reapir ansi colors in cmd.exe
func Init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
