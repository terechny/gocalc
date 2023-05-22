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
	fmt.Print("Введите строку: ")
	text, _ := reader.ReadString('\n')

	text = strings.TrimRight(text, "\r\n")

	parts := strings.Split(text, " ")

	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный формат ввода")
		return
	}

	// Парсим операнды
	a, err := parseOperand(parts[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := parseOperand(parts[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Выполняем операцию
	var result int
	switch parts[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неверный оператор")
		return
	}

	// Выводим результат
	if isRoman(text) {
		if result <= 0 {
			fmt.Println("Ошибка: результат должен быть положительным")
			return
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}

}

// Проверяем, является ли строка римским числом
func isRoman(input string) bool {
	return strings.ContainsAny(input, "IVXLC")
}

// Парсим операнд
func parseOperand(operand string) (int, error) {
	if isRoman(operand) {
		return romanToInt(operand)
	} else {
		num, err := strconv.Atoi(operand)
		if err != nil || num < 1 || num > 10 {
			return 0, fmt.Errorf("Ошибка: неверный операнд %s", err)
		}
		return num, nil
	}
}

// Преобразуем римское число в арабское
func romanToInt(roman string) (int, error) {
	var result int

	for i := 0; i < len(roman); i++ {

		if i > 0 && romanToIntMap[roman[i]] > romanToIntMap[roman[i-1]] {

			result += romanToIntMap[roman[i]] - 2*romanToIntMap[roman[i-1]]

		} else {

			result += romanToIntMap[roman[i]]
		}
	}

	return result, nil

}

// Преобразуем арабское число в римское
func intToRoman(arabic int) string {

	if arabic < 1 || arabic > 3999 {
		return "Неверное число"
	}
	var result string
	for arabic >= 1000 {
		result += "M"
		arabic -= 1000
	}
	if arabic >= 900 {
		result += "CM"
		arabic -= 900
	}
	if arabic >= 500 {
		result += "D"
		arabic -= 500
	}
	if arabic >= 400 {
		result += "CD"
		arabic -= 400
	}
	for arabic >= 100 {
		result += "C"
		arabic -= 100
	}
	if arabic >= 90 {
		result += "XC"
		arabic -= 90
	}
	if arabic >= 50 {
		result += "L"
		arabic -= 50
	}
	if arabic >= 40 {
		result += "XL"
		arabic -= 40
	}
	for arabic >= 10 {
		result += "X"
		arabic -= 10
	}
	if arabic >= 9 {
		result += "IX"
		arabic -= 9
	}
	if arabic >= 5 {
		result += "V"
		arabic -= 5
	}
	if arabic >= 4 {
		result += "IV"
		arabic -= 4
	}
	for arabic > 0 {
		result += "I"
		arabic--
	}
	return result

}

// Карта для преобразования римских чисел в арабские
var romanToIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}
