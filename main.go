package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Number: struct and method

type Number struct {
	value float64
}

func (n Number) calculate() float64 {
	return n.value
}

// Addition: struct and method

type Addition struct {
	leftNumber  float64
	rightNumber float64
}

func (a Addition) calculate() float64 {
	return a.leftNumber + a.rightNumber
}

// Subtraction: struct and method

type Subtraction struct {
	leftNumber  float64
	rightNumber float64
}

func (s Subtraction) calculate() float64 {
	return s.leftNumber - s.rightNumber
}

// Multiplication: struct and method

type Multiplication struct {
	leftNumber  float64
	rightNumber float64
}

func (m Multiplication) calculate() float64 {
	return m.leftNumber * m.rightNumber
}

// Division: struct and method

type Division struct {
	leftNumber  float64
	rightNumber float64
}

func (d Division) calculate() float64 {
	return d.leftNumber / d.rightNumber
}

// Expression: interface of methods

type Expression interface {
	calculate() float64
}

// function which handles only four valid operations

func operation(operator string, leftNum float64, rightNum float64) (Expression, error) {
	switch operator {
	case "+":
		return Addition{leftNumber: leftNum, rightNumber: rightNum}, nil
	case "-":
		return Subtraction{leftNumber: leftNum, rightNumber: rightNum}, nil
	case "*":
		return Multiplication{leftNumber: leftNum, rightNumber: rightNum}, nil
	case "/":
		return Division{leftNumber: leftNum, rightNumber: rightNum}, nil
	default:
		return nil, errors.New("invalid operator")
	}
}

// funcion which checks input validation using parse tree

func parseTree(str string) (Expression, error) {

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

	switch len(tokens) {

	// the number of token is 0
	case 0:
		return nil, errors.New("input is empty")

	// the number of token is 1
	case 1:
		number, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			return nil, errors.New("invalid number")
		}

		return Number{value: number}, nil

	// the number of token is 2
	case 2:
		return nil, errors.New("insufficient number of numbers or operator")

	// the number of token is 3
	case 3:
		// get the first token
		leftNumber, err := strconv.ParseFloat(tokens[0], 64)
		// check if the first token is valid
		if err != nil {
			return nil, errors.New("invalid left number")
		}

		// get the third token
		rightNumber, err := strconv.ParseFloat(tokens[2], 64)
		// check if the third token is valid
		if err != nil {
			return nil, errors.New("invalid right number")
		}

		// get the corressponding operation
		expression, err := operation(tokens[1], leftNumber, rightNumber)
		// check if the second token is valid
		if err != nil {
			return nil, err
		}

		return expression, nil

	// the number of token is more than 3
	default:
		return nil, errors.New("too many numbers or operators")
	}
}

// main function

func main() {

	// main loop
	for true {

		// input from command line
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// conditions of loop termination
		if input == "exit" || input == "Exit" || input == "EXIT" {
			break
		}

		// make a parse tree using input
		equation, err := parseTree(input)

		// if an error exists in the parse tree, the corresponding error will be printed
		if err != nil {
			fmt.Println("Error:", err)
			// otherwise, the calculated value will be printed
		} else {
			fmt.Println(equation.calculate())
		}

	}
}
