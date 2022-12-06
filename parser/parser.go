package parser

import (
	"calculator_project/operation"
	"errors"
	"strconv"
	"strings"
)

// GetTokens: given a string, get the array of all tokens

func GetTokens(str string) []string {
	chars := []rune(str)
	var token string
	var tokens []string
	var operators = "+-*/"

	// iterate over characters in argument str
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])

		// when char is one of valid operators
		if strings.Contains(operators, char) {
			if len(token) > 0 && len(tokens) == 0 {
				tokens = append(tokens, token)
				token = ""
				tokens = append(tokens, char)
			} else if len(tokens) == 1 {
				tokens = append(tokens, char)
			} else {
				token += char
			}

			// when char is space
		} else if char == " " {
			if len(token) > 0 {
				tokens = append(tokens, token)
				token = ""
			}

			// when char is any other character
		} else {
			token += char
		}
	}

	// append the last token to tokens if it is non-empty
	if len(token) > 0 {
		tokens = append(tokens, token)
	}

	return tokens
}

// GetExpressionWithOneToken: given one token, get the corresponding expression if the token is valid

func GetExpressionWithOneToken(token string) (operation.Expression, error) {
	number, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return nil, errors.New("invalid number")
	}

	return operation.Number{Value: number}, nil
}

// GetExpressionWithThreeTokens: given three tokens, get the corresponding expression if all the tokens are valid

func GetExpressionWithThreeTokens(firstToken string, secondToken string, thridToken string) (operation.Expression, error) {
	// check if the first token is valid
	leftNumber, err := strconv.ParseFloat(firstToken, 64)
	if err != nil {
		return nil, errors.New("invalid left number")
	}

	// check if the third token is valid
	rightNumber, err := strconv.ParseFloat(thridToken, 64)
	if err != nil {
		return nil, errors.New("invalid right number")
	}

	// check if the second token is valid
	expression, err := operation.GetExpression(secondToken, leftNumber, rightNumber)
	if err != nil {
		return nil, err
	}

	return expression, nil
}

// Parser: given an array of tokens, get the corresponding valid Expression

func Parser(tokens []string) (operation.Expression, error) {
	switch len(tokens) {
	case 0:
		return nil, errors.New("input is empty")
	case 1:
		return GetExpressionWithOneToken(tokens[0])
	case 2:
		return nil, errors.New("insufficient number of numbers or operator")
	case 3:
		return GetExpressionWithThreeTokens(tokens[0], tokens[1], tokens[2])
	default:
		return nil, errors.New("too many numbers or operators")
	}
}
