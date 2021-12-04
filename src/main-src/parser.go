package main

import (
	"fmt"
	"strings"
)

func getLastTokenAST(actions []Action) Action {
	returntoken := createAction(NIL, []string{NIL, NIL})

	if len(actions) == 0 {
		returntoken = createAction(NIL, []string{NIL, NIL})
	} else {
		returntoken = actions[len(actions)-1]
	}

	return returntoken
}

var (
	tokenAST                        []Action
	functionBracketsExpectation     int = 0
	functionNameExpectation         int = 0
	functionCloseBracketExpectation int = 0
)

func parser(token []Token) {

	for i := 0; i < len(token); i++ {
		var nextToken Token = createToken(NIL, NIL, 1)
		if len(token)+1 > i {
			if len(token)-1 == i {
				nextToken = createToken(NIL, NIL, 1)
			} else {
				nextToken = token[i+1]
			}
		}
		currentToken := token[i]
		lastTokenAST := getLastTokenAST(tokenAST)
		Use(lastTokenAST)

		if functionCloseBracketExpectation == 1 {
			// if currentToken.Type == CLOSEBRACKET
		} else if functionNameExpectation == 1 {
			if currentToken.Type == IDENTIFIER {
				functionBracketsExpectation = 1
				functionNameExpectation = 0
			} else {
				ece("Name of function must be type "+IDENTIFIER+" not "+currentToken.Type, PARSER_ERROR_TITLE)
			}
		} else if functionBracketsExpectation == 1 {
			fmt.Println(currentToken.Type)
			fmt.Println(currentToken.Value)
			if currentToken.Value == OPENBRACKET {
				functionBracketsExpectation = 0
				functionCloseBracketExpectation = 1
			} else {
				ece("Expected "+OPENPARENT+" but found "+currentToken.Type, PARSER_ERROR_TITLE)
			}
		} else if currentToken.Type == INT {

		} else if currentToken.Type == IDENTIFIER {
			if currentToken.Value == KEYWORD_FUNCTION {
				functionNameExpectation = 1
			} else if currentToken.Value == KEYWORD_RETURN {
				// Return
			} else if currentToken.Value == KEYWORD_VARIABLE {
				// Var
			} else if currentToken.Value == KEYWORD_CONSTANT {
				// Const
			} else if currentToken.Value == KEYWORD_TRUE {
				// True
			} else if currentToken.Value == KEYWORD_FALSE {
				// False
			} else if currentToken.Value == KEYWORD_IF {
				// If
			} else if currentToken.Value == KEYWORD_ELSE {
				// Else
			} else if currentToken.Value == KEYWORD_IMPORT {
				// Import packages
			} else if currentToken.Value == KEYWORD_INPCODE {
				// Input code in go
			} else {
				// Is it a variable?
			}
		} else if currentToken.Type == STRING {

		} else {

		}
		Use(nextToken)
	}

	if functionNameExpectation == 1 {
		ece("Expected function name!", PARSER_ERROR_TITLE)
	}

	for i := 0; i < len(tokenAST); i++ {
		fmt.Println(tokenAST[i].actionType + ":" + colorYellow + strings.Join(tokenAST[i].body, " "))
	}
}
