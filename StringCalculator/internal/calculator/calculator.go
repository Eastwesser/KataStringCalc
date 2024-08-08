package calculator

import (
	"KataStringCalc/internal/calculator/add"
	"KataStringCalc/internal/calculator/divide"
	"KataStringCalc/internal/calculator/multiply"
	"KataStringCalc/internal/calculator/subtract"
	"strings"
)

func Calculate(input string) string {
	input = strings.TrimSpace(input)
	parts := strings.SplitN(input, " ", 3)

	if len(parts) != 3 {
		panic("Некорректное выражение")
	}

	str1 := parts[0]
	operator := parts[1]
	str2 := parts[2]

	str1 = strings.Trim(str1, "\"")
	str2 = strings.Trim(str2, "\"")

	if len(str1) == 0 || len(str2) == 0 {
		panic("Некорректное выражение")
	}

	switch operator {

	case "+":
		return add.Add(str1, str2)

	case "-":
		return subtract.Subtract(str1, str2)

	case "*":
		return multiply.Multiply(str1, str2)

	case "/":
		return divide.Divide(str1, str2)

	default:
		panic("Недопустимая операция")
	}
}
