GO = go
GCC = gcc
GO_SRC = src/main-src/main.go
GCC_SRC = src/first-src/main.c

compile_go:
	@echo Compiling Interpreter (Go)
	@cd src/main-src/
	@go build