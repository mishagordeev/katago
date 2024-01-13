package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значение")
	input, _ := reader.ReadString('\n')
	fmt.Println(calc(input))
}

func calc(input string) string {
	var result string

	for {
		input = strings.TrimSpace(input)
		inputArray := strings.Split(input, " ")

		if !isInputLengthCorrect(inputArray) {
			result = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор."
			return result
		}

		operator := inputArray[1]

		if !isOperatorCorrect(operator) {
			result = "Вывод ошибки, так как оператор не удовлетворяет заданию (+, -, *, /)."
			return result
		}

		type Operand struct {
			value         string
			numeralSystem string
		}

		a := Operand{value: inputArray[0]}
		b := Operand{value: inputArray[2]}

		if !isOperandCorrect(a.value) {
			result = "Вывод ошибки, так как операнды не удовлетворяет заданию (целые арабские или римские цифры от 1 до 10)"
			return result
		}

		if !isOperandCorrect(b.value) {
			result = "Вывод ошибки, так как операнды не удовлетворяет заданию (целые арабские или римские цифры от 1 до 10)"
			return result
		}

		a.numeralSystem = defineNumeralSystem(a.value)
		b.numeralSystem = defineNumeralSystem(b.value)

		if !isNumeralSystemCorrect(a.numeralSystem, b.numeralSystem) {
			result = "Вывод ошибки, так как используются одновременно разные системы счисления."
			return result
		}

		result = calculate(a.value, b.value, operator, a.numeralSystem)
		return result
	}
}

func isInputLengthCorrect(inputArray []string) bool {
	var result bool
	if len(inputArray) == 3 {
		result = true
	} else {
		result = false
	}
	return result
}

func isOperatorCorrect(operator string) bool {
	var result bool
	correctOperators := []string{"+", "-", "*", "/"}
	if slices.Contains(correctOperators, operator) {
		result = true
	} else {
		result = false
	}
	return result
}

func isOperandCorrect(operand string) bool {
	var result bool
	correctRomanOperands := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	integer, err := strconv.Atoi(operand)
	if err != nil {
		if slices.Contains(correctRomanOperands, operand) {
			result = true
		}
	} else {
		if integer >= 1 && integer < 10 {
			result = true
		}
	}
	return result
}

func defineNumeralSystem(operand string) string {
	_, err := strconv.Atoi(operand)
	if err != nil {
		return "roman"
	} else {
		return "arabic"
	}
}

func isNumeralSystemCorrect(operand1, operand2 string) bool {
	var result bool
	if strings.Compare(operand1, operand2) == 0 {
		result = true
	} else {
		result = false
	}
	return result
}

func calculate(operand1, operand2, operator, numeralSystem string) string {
	var result string
	var integerResult int

	if strings.Compare(numeralSystem, "arabic") == 0 {
		integerOperand1, _ := strconv.Atoi(operand1)
		integerOperand2, _ := strconv.Atoi(operand2)

		switch operator {
		case "+":
			integerResult = integerOperand1 + integerOperand2
		case "-":
			integerResult = integerOperand1 - integerOperand2
		case "*":
			integerResult = integerOperand1 * integerOperand2
		case "/":
			integerResult = integerOperand1 / integerOperand2
		}
		result = strconv.Itoa(integerResult)
	}

	if strings.Compare(numeralSystem, "roman") == 0 {
		romanNumerals := map[string]int{
			"I":    1,
			"II":   2,
			"III":  3,
			"IV":   4,
			"V":    5,
			"VI":   6,
			"VII":  7,
			"VIII": 8,
			"IX":   9,
			"X":    10,
		}

		switch operator {
		case "+":
			integerResult = romanNumerals[operand1] + romanNumerals[operand2]
		case "-":
			integerResult = romanNumerals[operand1] - romanNumerals[operand2]
		case "*":
			integerResult = romanNumerals[operand1] * romanNumerals[operand2]
		case "/":
			integerResult = romanNumerals[operand1] / romanNumerals[operand2]
		}

		if integerResult < 1 {
			result = "Вывод ошибки, так как в римской системе нет отрицательных чисел и нуля."
			return result
		}

		romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		result = ""

		for i := 0; i < len(romanSymbols); i++ {
			for integerResult >= romanValues[i] {
				result += romanSymbols[i]
				integerResult -= romanValues[i]
			}
		}
	}
	return result
}
