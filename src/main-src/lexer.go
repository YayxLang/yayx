package main

import (
	"fmt"
	"strings"
)

func getLastToken(token []Token) Token {
	returntoken := createToken(NIL, NIL, 0)

	if len(token) == 0 {
		returntoken = createToken(NIL, NIL, 0)
	} else {
		returntoken = token[len(token)-1]
	}

	return returntoken
}

// Lexer function
func lexer(value string) []Token {
	var currentChar string = ""
	var token []Token
	var space int = 0
	var stringOpened int = 0
	var splitedString = strings.Split(value, "")
	var commentOpended int = 0

	for i := 0; i < len(splitedString); i++ {
		currentChar = splitedString[i]
		lastToken := getLastToken(token)

		if currentChar == "#" {
			if commentOpended == 1 {
				commentOpended = 0
			} else {
				commentOpended = 1
			}
		} else if commentOpended == 1 {
			continue
		} else if currentChar == `"` {
			if stringOpened == 1 {
				stringOpened = 0
			} else {
				token = append(token, createToken("", STRING, 0))
				stringOpened = 1
			}
		} else if currentChar == " " {
			if stringOpened == 1 {
				token[len(token)-1].Value += currentChar
			} else {
				space = 1
			}
		} else if stringOpened == 1 {
			if lastToken.Type == NIL {
				token = append(token, createToken(currentChar, STRING, i))
			} else {
				token[len(token)-1].Value += currentChar
			}
		} else if strings.ContainsAny(CONSTANTS_LETTER, currentChar) {
			if lastToken.Type == NIL {
				token = append(token, createToken(currentChar, IDENTIFIER, i))
				space = 0
			} else {
				// Space working
				//fmt.Println(strconv.Itoa(space) + " | " + colorPurple + currentChar + colorReset)
				if space == 0 {
					if lastToken.Type == IDENTIFIER {
						token[len(token)-1].Value += currentChar
						space = 0
					} else {
						token = append(token, createToken(currentChar, IDENTIFIER, i))
						space = 0
					}
				} else {
					token = append(token, createToken(currentChar, IDENTIFIER, i))
					space = 0
				}
			}
		} else if strings.ContainsAny(CONSTANTS_NUMBER, currentChar) {
			if lastToken.Type == NIL {
				token = append(token, createToken(currentChar, INT, i))
			} else {
				if token[len(token)-1].Type == INT || token[len(token)-1].Type == FLOAT {
					if token[len(token)-1].Type == INT {
						token[len(token)-1].Value += currentChar
					} else {
						token[len(token)-1].Type = FLOAT
						token[len(token)-1].Value += currentChar
					}
				} else {
					token = append(token, createToken(currentChar, INT, i))
				}
			}
		} else if currentChar == PLUS {
			token = append(token, createToken(PLUS, TOKENNAME_PLUS, i))
		} else if currentChar == COLON {
			token = append(token, createToken(COLON, TOKENNAME_COLON, i))
		} else if currentChar == MINUS {
			token = append(token, createToken(MINUS, TOKENNAME_MINUS, i))
		} else if currentChar == MUL {
			token = append(token, createToken(MUL, TOKENNAME_MUL, i))
		} else if currentChar == DIV {
			token = append(token, createToken(DIV, TOKENNAME_DIV, i))
		} else if currentChar == BANG {
			token = append(token, createToken(BANG, TOKENNAME_BANG, i))
		} else if currentChar == GT {
			token = append(token, createToken(GT, TOKENNAME_GT, i))
		} else if currentChar == LT {
			token = append(token, createToken(LT, TOKENNAME_LT, i))
		} else if currentChar == ASN {
			token = append(token, createToken(ASN, TOKENNAME_ASN, i))
		} else if currentChar == COMMA {
			token = append(token, createToken(COMMA, TOKENNAME_COMMA, i))
		} else if currentChar == DOT {
			if lastToken.Type == NIL {
				token = append(token, createToken(currentChar, INT, i))
			} else {
				if token[len(token)-1].Type == INT || token[len(token)-1].Type == FLOAT {
					if token[len(token)-1].Type == INT {
						token[len(token)-1].Type = FLOAT
						token[len(token)-1].Value += currentChar
					} else {
						if strings.ContainsAny(token[len(token)-1].Value, ".") == true {
							ece(ERROR_LEXER_FLOAT_DOT, LEXER_ERROR_TITLE)
						}
						token[len(token)-1].Type = FLOAT
						token[len(token)-1].Value += currentChar
					}
				} else {
					token = append(token, createToken("0"+currentChar, INT, i))
				}
			}
		} else if currentChar == OPENPARENT {
			token = append(token, createToken(OPENPARENT, TOKENNAME_OPENPARENT, i))
		} else if currentChar == CLOSEPARENT {
			token = append(token, createToken(CLOSEPARENT, TOKENNAME_CLOSEPARENT, i))
		} else if currentChar == OPENBRACKET {
			token = append(token, createToken(OPENBRACKET, TOKENNAME_OPENBRACKET, i))
		} else if currentChar == CLOSEBRACKET {
			token = append(token, createToken(CLOSEBRACKET, TOKENNAME_CLOSEBRACKET, i))
		} else if currentChar == OPENBRACE {
			token = append(token, createToken(OPENBRACE, TOKENNAME_OPENBRACE, i))
		} else if currentChar == CLOSEBRACE {
			token = append(token, createToken(CLOSEBRACE, TOKENNAME_CLOSEBRACE, i))
		}
	}

	for i := 0; i < len(token); i++ {
		fmt.Println(token[i].Type + ":" + colorYellow + " " + token[i].Value + colorReset)
	}

	return token
}
