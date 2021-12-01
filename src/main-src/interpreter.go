package main

func interpreter_run(value string) {
	token := lexer(value)

	parser(token)
}
