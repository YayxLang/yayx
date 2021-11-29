package main

func interpreter_run(filename string) {
	value := read_file(filename)

	lexer(value)
}
