package main

var (
	tokenAST []Token
)

func parser(token []Token) {
	for i := 0; i < len(token); i++ {
		currentToken := token[i]

		if currentToken.Type == INT {

		} else if currentToken.Type == IDENTIFIER {
			if currentToken.Value == 
		} else if currentToken.Type == STRING {

		} else {

		}
	}

	Use(tokenAST)
}
