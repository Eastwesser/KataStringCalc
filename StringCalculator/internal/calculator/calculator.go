package calculator

import (
	"strconv"
	"strings"
)

func Calculate(input string) string {
	input = strings.TrimSpace(input) // Удаляем пробелы

	parts := strings.Fields(input) // Разделяем строку на части

	// Если операндов больше двух, выводим панику
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

// add Складывает две строки и вызывает truncate.
func add(a, b string) string {
	return truncate(a + b)
}

// subtract Удаляет все вхождения второй строки из первой и вызывает truncate.
func subtract(a, b string) string {
	if !strings.Contains(a, b) {
		return truncate(a)
	}
	return truncate(strings.ReplaceAll(a, b, ""))
}

// multiply Повторяет первую строку указанное количество раз и вызывает truncate.
func multiply(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для умножения")
	}
	return truncate(strings.Repeat(a, n))
}

// divide Делит первую строку на указанное количество частей и вызывает truncate.
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

// Если строка длиннее 40 символов, truncate обрезает её до 40 символов и добавляет '...'
func truncate(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}
