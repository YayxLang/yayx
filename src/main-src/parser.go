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

func parser(token []Token) {

	var tokenAST []Action
	var mathematicalOperationOpen int = 0

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
		if currentToken.Type == INT {
			if mathematicalOperationOpen == 1 {
				if currentToken.Type == INT || currentToken.Type == PLUS || currentToken.Type == MINUS || currentToken.Type == MUL || currentToken.Type == DIV {
					if lastTokenAST.actionType == "Math" {
						tokenAST[len(tokenAST)-1].body = append(tokenAST[len(tokenAST)-1].body, currentToken.Value)
					}
				}
			} else {
				tokenAST = append(tokenAST, createAction("Math", []string{currentToken.Value}))
				mathematicalOperationOpen = 1
			}
		} else if currentToken.Type == IDENTIFIER {

		} else if currentToken.Type == STRING {

		} else {
			if mathematicalOperationOpen == 1 {
				if currentToken.Type == INT || currentToken.Type == PLUS || currentToken.Type == MINUS || currentToken.Type == MUL || currentToken.Type == DIV {
					if lastTokenAST.actionType == "Math" {
						tokenAST[len(tokenAST)-1].body = append(tokenAST[len(tokenAST)-1].body, currentToken.Value)
					}
				}
			} else {
				tokenAST = append(tokenAST, createAction("Math", []string{currentToken.Value}))
				mathematicalOperationOpen = 1
			}
		}
		Use(nextToken)
	}

	for i := 0; i < len(tokenAST); i++ {
		fmt.Println(tokenAST[i].actionType + ":" + colorYellow + strings.Join(tokenAST[i].body, " "))
	}
}
