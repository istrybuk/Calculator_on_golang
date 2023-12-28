package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const alphaRoman = "IVXLCDM"

func ArabicToRoman(number int) (string, error) {
	// Преобразование числа в римское значение
	if number < 1 {
		return "", errors.New("неверный расчет: результат меньше 1")
	}

	conversions := []struct {
		value   int
		numeral string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.numeral
			number -= conversion.value
		}
	}
	return roman, nil
}

func DecodeToArabic(roman string) int {
	// Преобразование чисел
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var decNum, tmpNum int
	for i := len(roman) - 1; i >= 0; i-- {
		romanDigit := roman[i]
		decDigit := translateRoman[romanDigit]
		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
	}
	return decNum
}

func CheckString(input1, input2 string) (interface{}, error) {
	// Проверяем, содержатся ли оба аргумента в диапазоне от 0 до 10
	num1, err1 := strconv.Atoi(input1)
	num2, err2 := strconv.Atoi(input2)
	if err1 == nil && err2 == nil {
		if num1 >= 0 && num1 <= 10 && num2 >= 0 && num2 <= 10 {
			return []int{num1, num2}, nil
		}
		return nil, errors.New("введеное число должно находится в диапазоне от 0 до 10")
	}

	// Проверяем, содержат ли оба аргумента только символы 'I', 'V', 'X', 'L', 'C', 'D', 'M'
	if strings.IndexFunc(input1, func(r rune) bool {
		return !strings.ContainsRune(alphaRoman, r)
	}) == -1 && strings.IndexFunc(input2, func(r rune) bool {
		return !strings.ContainsRune(alphaRoman, r)
	}) == -1 {
		return []string{input1, input2}, nil
	}

	if err1 == nil || err2 == nil {
		return nil, errors.New("недопустимо использовать одновременно цифры и римской буквы")
	}

	return nil, errors.New("введены некорректные значения, введите 'I', 'V', 'X', 'L', 'C', 'D', 'M'")
}

func Calculate(x int, operator string, y int) int {
	var result int
	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			fmt.Println("Division by zero")
			os.Exit(1)
		}
		result = x / y
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	expression = strings.TrimSpace(expression) // Убирает пробелы в начале и в конце строки

	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Строка не является математической операцией.")
		os.Exit(1)
	}

	x, operator, y := tokens[0], tokens[1], tokens[2]

	result, err := CheckString(x, y)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		switch res := result.(type) {
		case []int: // если аргументы цифры
			num1, num2 := res[0], res[1]

			resultArabic := Calculate(num1, operator, num2)
			fmt.Println("Результат вырожения ", resultArabic)
		case []string: // если аргументы строки
			str1, str2 := res[0], res[1]
			decode1, decode2 := DecodeToArabic(str1), DecodeToArabic(str2)
			calcRoman := Calculate(decode1, operator, decode2)
			resultRoman, err := ArabicToRoman(calcRoman)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Результат вырожения ", resultRoman)
			}
		}
	}
}
