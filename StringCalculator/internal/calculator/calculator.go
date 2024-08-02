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

// add Складывает две строки и вызывает trimStringsAfter40.
func add(a, b string) string {
	return trimStringsAfter40(a + b)
}

// subtract Удаляет все вхождения второй строки из первой и вызывает trimStringsAfter40.
func subtract(a, b string) string {
	if !strings.Contains(a, b) {
		return trimStringsAfter40(a)
	}
	return trimStringsAfter40(strings.ReplaceAll(a, b, ""))
}

// multiply Повторяет первую строку указанное количество раз и вызывает trimStringsAfter40.
func multiply(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для умножения")
	}
	return trimStringsAfter40(strings.Repeat(a, n))
}

// divide Делит первую строку на указанное количество частей и вызывает trimStringsAfter40.
func divide(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil || n < 1 || n > 10 {
		panic("Некорректное число для деления")
	}
	if len(a) < n {
		return ""
	}
	return trimStringsAfter40(a[:len(a)/n])
}

// Если строка длиннее 40 символов, trimStringsAfter40 обрезает её до 40 символов и добавляет '...'
func trimStringsAfter40(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}
