package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// maps для конвертации систем счисления 
var romeToArab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabToRome = map[int]string{
	100: "C", 90: "XC", 50: "L", 40: "XL",
	10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
}

// I/O
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение: ")

	for scanner.Scan() {
		input := scanner.Text()
		result := calc(input)
		fmt.Println("Результат:", result)
		fmt.Print("Введите выражение: ")
	}
}

// основная функция
func calc(input string) string {
	parts := strings.Fields(input)
	// проверка количества смиволов
	if len(parts) == 1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	} else if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	num1, op, num2 := parts[0], parts[1], parts[2]
	num1, num2 = checkInteger(num1), checkInteger(num2)

	// проверка, введены римские или арабские числа
	checkRome := checkRomeNum(num1) && checkRomeNum(num2)
	checkArab := checkArabNum(num1) && checkArabNum(num2)

	var num1Int, num2Int int
	if checkRome {
		// конвертаця римских в арабские
		num1Int = romeToArab[num1]
		num2Int = romeToArab[num2]
	// проверка наличия чисел в обеих maps
	} else if !(checkRomeNum(num1) || checkArabNum(num1)) || !(checkRomeNum(num2) || checkArabNum(num2)) {
		panic("Выдача паники, так как строка не является математической операцией.")
	} else if !checkRome && !checkArab {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else {
		// конвертация арабских в int
		num1Int, _ = strconv.Atoi(num1)
		num2Int, _ = strconv.Atoi(num2)
	}

	// проверка диапазона
	if num1Int < 1 || num1Int > 10 || num2Int < 1 || num2Int > 10 {
		panic("Выдача паники, так как принимаются числа только от 1 до 10 включительно")
	}
	
	// получение результата операции
	var result int
	switch op {
	case "+":
		result = num1Int + num2Int
	case "-":
		result = num1Int - num2Int
	case "*":
		result = num1Int * num2Int
	case "/":
		result = num1Int / num2Int
	default:
		panic("Выдача паники, так как операнд некорректен (+, -, /, *).")
	}

	// проверка на римские меньше 1
	if checkRome {
		if result == 0 {
			panic("Выдача паники, так как в римской системе нет нуля.")
		} else if result < 0 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		return getArabToRome(result)
	}

	// возвращение результата в арабских
	return strconv.Itoa(result)
}

// функция перевода арабских чисел в римские
func getArabToRome(number int) string {
	var result strings.Builder
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, value := range values {
		// сравнение числа со списком римских, хз как через саму map сделать
		for number >= value {
			result.WriteString(arabToRome[value])
			number -= value
		}
	}
	return result.String()
}

// проверка на наличие в словаре римских чисел
func checkRomeNum(s string) bool {
	_, value := romeToArab[s]
	return value
}

// проверка на наличие в словаре арабских чисел
func checkArabNum(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

// Проверка на целое число
func checkInteger(s string) string {
	if strings.Contains(s, ".") {
		parts := strings.Split(s, ".")
		// Проверка дробного числа на возможность удалить все нули после точки
		if len(parts) == 2 && strings.Trim(parts[1], "0") == "" {
			s = parts[0]
		} else {
			panic("Выдача паники, так как принимаются только целые числа.")
		}
	}
	return s
}
