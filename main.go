package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
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
	if isRomanNumeral(a) && isRomanNumeral(b) {
		// Если оба операнда являются римскими числами, то результат тоже должен быть римским числом
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Ошибка при преобразовании в римское число:", err)
			return
		}
		fmt.Println(romanResult)
	} else {
		// Иначе выводим арабское число
		fmt.Println(result)
	}
}

func isRomanNumeral(input string) bool {
	// Проверка, является ли строка римским числом
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var prevValue int
	var consecutiveFailures int
	for _, char := range input {
		value, ok := romanNumerals[char]
		if !ok {
			return false
		}
		if value > prevValue {
			consecutiveFailures = 0
		} else if value < prevValue {
			consecutiveFailures++
			if consecutiveFailures > 1 {
				return false // Римское число неверно сформировано
			}
		}
		prevValue = value
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
	if num < 1 || num > 10 {
		return 0, errors.New("число вне допустимого диапазона (1-10)")
	}
	return num, nil
}

func calculate(a, b int, operator string) (int, error) {
	// Выполнение операции
	switch operator {
	case "+":
		result := a + b
		if isRomanNumeral(strconv.Itoa(a)) && isRomanNumeral(strconv.Itoa(b)) {
			if result < 1 || result > 3999 {
				return 0, errors.New("результат вне допустимого диапазона (1-3999)")
			}
		}
		return result, nil
	case "-":
		result := a - b
		if result < 1 {
			return 0, errors.New("результат меньше 1 (римские числа не могут быть отрицательными)")
		}
		return result, nil
	case "*":
		result := a * b
		if isRomanNumeral(strconv.Itoa(a)) && isRomanNumeral(strconv.Itoa(b)) {
			if result < 1 || result > 3999 {
				return 0, errors.New("результат вне допустимого диапазона (1-3999)")
			}
		}
		return result, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		result := a / b
		if isRomanNumeral(strconv.Itoa(a)) && isRomanNumeral(strconv.Itoa(b)) {
			if result < 1 {
				return 0, errors.New("результат меньше 1 (римские числа не могут быть отрицательными)")
			}
		}
		return result, nil
	default:
		return 0, errors.New("недопустимая операция")
	}
}

func arabicToRoman(arabic int) (string, error) {
	// Преобразование арабских чисел в римские
	// Добавьте сюда свою реализацию
	return "", errors.New("преобразование арабских чисел в римские числа пока не реализовано")
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
		value := romanNumerals[rune(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result, nil
}
