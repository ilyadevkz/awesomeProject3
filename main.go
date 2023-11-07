package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите математическое выражение (например, '1 + 2' или 'VI / III):")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем введенную строку на операнды и оператор
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Ошибка: формат математической операции не удовлетворяет заданию")
		return
	}

	// Получаем операнды и оператор
	a := parts[0]
	operator := parts[1]
	b := parts[2]

	// Проверка на использование разных систем счисления
	if (isRomanNumeral(a) && isArabicNumeral(b)) || (isArabicNumeral(a) && isRomanNumeral(b)) {
		fmt.Println("Ошибка: используются одновременно разные системы счисления.")
		return
	}

	// Преобразуем операнды в числа
	numA, errA := convertToNumber(a)
	numB, errB := convertToNumber(b)

	if errA != nil || errB != nil {
		fmt.Println("Ошибка: введены недопустимые числа.")
		return
	}

	// Выполняем операцию
	result, err := calculate(numA, numB, operator)
	if err != nil {
		fmt.Println("Ошибка: недопустимая операция.")
		return
	}

	// Выводим результат
	fmt.Println(result)
}

func isRomanNumeral(input string) bool {
	// Проверка, является ли строка римским числом
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		// Добавьте другие римские цифры по мере необходимости
	}
	for _, char := range input {
		if _, ok := romanNumerals[char]; !ok {
			return false
		}
	}
	return true
}

func isArabicNumeral(input string) bool {
	// Проверка, является ли строка арабским числом
	_, err := strconv.Atoi(input)
	return err == nil
}

func convertToNumber(input string) (int, error) {
	// Преобразование строки в число
	if isRomanNumeral(input) {
		return romanToArabic(input)
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func calculate(a, b int, operator string) (int, error) {
	// Выполнение операции
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("недопустимая операция")
	}
}

func romanToArabic(roman string) (int, error) {
	// Реализация преобразования римских чисел в арабские
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	prevValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i)]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result, nil
}

