package calculator

import (
	"strconv"
	"strings"
)

func Calculate(input string) string {
	// Удаляем пробелы
	input = strings.TrimSpace(input)

	// Разделяем строку на части
	parts := strings.Fields(input)

	if len(parts) != 3 {
		panic("Некорректное выражение")
	}

	// Получаем операнды и оператор
	str1 := parts[0]
	operator := parts[1]
	str2 := parts[2]

	// Убираем кавычки
	str1 = strings.Trim(str1, "\"")
	str2 = strings.Trim(str2, "\"")

	switch operator {
	case "+":
		return add(str1, str2)
	case "-":
		return subtract(str1, str2)
	case "*":
		return multiply(str1, parts[2])
	case "/":
		return divide(str1, parts[2])
	default:
		panic("Недопустимая операция")
	}
}

func add(a, b string) string {
	return truncate(a + b)
}

func subtract(a, b string) string {
	if !strings.Contains(a, b) {
		return truncate(a)
	}
	return truncate(strings.ReplaceAll(a, b, ""))
}

func multiply(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для умножения")
	}
	return truncate(strings.Repeat(a, n))
}

func divide(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для деления")
	}
	if len(a) < n {
		return ""
	}
	return truncate(a[:len(a)/n])
}

func truncate(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}
