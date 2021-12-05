package main

func interpreter_run(value string) {
	token := lexer(value)
	Use(token)
	// parser(token)
}
