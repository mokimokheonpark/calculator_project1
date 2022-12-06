package main

import (
	"bufio"
	"calculator_project/parser"
	"fmt"
	"os"
)

// main function

func main() {

	// main loop
	for {
		// input from command line
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// conditions of loop termination
		if input == "exit" || input == "Exit" || input == "EXIT" {
			break
		}

		// using the input, get tokens
		tokens := parser.GetTokens(input)

		// using the tokens, get the corresponding expression or error
		expression, parserError := parser.Parser(tokens)

		// if parser finds an error, the corresponding error will be printed
		if parserError != nil {
			fmt.Println("Error:", parserError)
			// otherwise, operation will be done
		} else {
			result, operationError := expression.Operate()
			// if an error is found during operation, the corresponding error will be printed
			if operationError != nil {
				fmt.Println("Error:", operationError)
				// otherwise, the operated value will be printed
			} else {
				fmt.Println(result)
			}
		}
	}
}
