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

	fmt.Println("Добро пожаловать в приложение Калькулятор!")
	fmt.Println("Введите выражение (например, 2 + 3):")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Выход из приложения Калькулятор...")
			break
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func evaluateExpression(input string) (float64, error) {
	input = strings.TrimSpace(input)
	tokens := strings.Fields(input)
	if len(tokens) != 3 {
		return 0, fmt.Errorf("Некорректное количество токенов в выражении")
	}

	operand1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Некорректный операнд: %s", tokens[0])
	}

	operand2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, fmt.Errorf("Некорректный операнд: %s", tokens[2])
	}

	// Выполнение арифметической операции
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
			return 0, fmt.Errorf("Деление на ноль")
		}
		result = operand1 / operand2
	default:
		return 0, fmt.Errorf("Некорректный оператор: %s", operator)
	}

	return result, nil
}




