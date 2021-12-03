package main

import (
	"fmt"
	"strconv"
)

func getLastTokenAST(actions []Action) Action {
	returntoken := createAction(NIL, ]NIL])

	if len(actions) == 0 {
		returntoken = createAction(NIL, NIL)
	} else {
		returntoken = actions[len(actions)-1]
	}

	return returntoken
}

func parser(token []Token) {
	// for i := 0; i < len(token); i++ {
	// 	fmt.Println(token[i].Type + ":" + colorYellow + " " + token[i].Value + colorReset)
	// }

	var tokenAST []Action
	var mathematicalOperationOpen int = 0

	for i := 0; i < len(token); i++ {
		var nextToken Token = createToken(NIL, NIL, 1)
		fmt.Println(strconv.Itoa(len(token)) + " | " + strconv.Itoa(i))
		if len(token)+1 > i {
			if len(token)-1 == i {
				nextToken = createToken(NIL, NIL, 1)
			} else {
				nextToken = token[i+1]
			}
		}
		currentToken := token[i]

		if currentToken.Type == INT {
			if mathematicalOperationOpen == 1 {

			} else {

			}
		} else if currentToken.Type == IDENTIFIER {

		} else if currentToken.Type == STRING {

		} else {

		}
		Use(nextToken)
	}

	Use(tokenAST)
}
