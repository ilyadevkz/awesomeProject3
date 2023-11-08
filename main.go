package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Calculator App!")
	fmt.Println("Enter an expression (e.g., 2 + 3):")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting the Calculator App...")
			break
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

func evaluateExpression(input string) (float64, error) {
	// Split the input into operands and operator
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		return 0, fmt.Errorf("Invalid expression")
	}

	// Parse the operands
	operand1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid operand: %s", tokens[0])
	}

	operand2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid operand: %s", tokens[2])
	}

	// Perform the arithmetic operation
	operator := tokens[1]
	var result float64

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		result = operand1 / operand2
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}

	return result, nil
}
