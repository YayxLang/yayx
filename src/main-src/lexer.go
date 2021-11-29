package main

import (
	"fmt"
	"strings"
)

// Lexer function
func lexer(value string) {
	var currentChar string = ""
	var token [] Token
	var space int = 0
	var stringOpened int = 0

	var splitedString = strings.Split(value, "")

	for i := 0; i < len(splitedString); i++ {
		currentChar = splitedString[i]

		if currentChar == `"` {

		} else if currentChar == " " {

		} else if stringOpened == 1 {

		} else if strings.ContainsAny(CONSTANTS_LETTER, currentChar) {

		} else if strings.ContainsAny(CONSTANTS_NUMBER, currentChar) {

		} else if currentChar == PLUS {
			token = append(token, createToken(TOKENNAME_PLUS, PLUS, i))
		} else if currentChar == MINUS {
			token = append(token, createToken(TOKENNAME_MINUS, MINUS, i))
		} else if currentChar == MUL {
			token = append(token, createToken(TOKENNAME_MUL, MUL, i))
		} else if currentChar == DIV {
			token = append(token, createToken(TOKENNAME_DIV, DIV, i))
		} else if currentChar == BANG {
			token = append(token, createToken(TOKENNAME_BANG, BANG, i))
		} else if currentChar == GT {
			token = append(token, createToken(TOKENNAME_GT, GT, i))
		} else if currentChar == LT {
			token = append(token, createToken(TOKENNAME_LT, LT, i))
		} else if currentChar == ASN {
			token = append(token, createToken(TOKENNAME_ASN, ASN, i))
		} else if currentChar == COMMA {
			token = append(token, createToken(TOKENNAME_COMMA, COMMA, i))
		} else if currentChar == DOT {
			lastToken := token[len(token) - 1]

			if lastToken.Type == INT || lastToken.Type == FLOAT {
				if lastToken.Type == INT {

				} else {
					if currentChar == DOT {
						if strings.ContainsAny(lastToken.Value, DOT) == true {
							ece(LEXER_ERROR_TITLE, ERROR_LEXER_FLOAT_DOT)
						}
					} else {
						lastToken.Value += currentChar
					}
				}
			} else {
				token = append(token, createToken(TOKENNAME_DOT, DOT, i))
			}
		} else if currentChar == SEMICOL {
			token = append(token, createToken(TOKENNAME_SEMICOL, SEMICOL, i))
		}
	}

	// Appending token struct to token array
	//token = append(token, createToken(STRING, "sa", 12))
	Use(space)
	Use(stringOpened)
	Use(token)
	Use(currentChar)

	i := 0;
	for i = 0; i < len(token); i++ {
		fmt.Println(colorYellow + token[i].Type + colorReset + ":" + colorPurple + token[i].Value + colorReset);
	}
}
