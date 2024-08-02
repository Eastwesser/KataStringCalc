package calculator

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/add"
	"github.com/Eastwesser/KataStringCalc/internal/calculator/divide"
	"github.com/Eastwesser/KataStringCalc/internal/calculator/multiply"
	"github.com/Eastwesser/KataStringCalc/internal/calculator/subtract"
	"strings"
)

// Calculate обрабатывает входную строку, разбивает её на части и выполняет соответствующую операцию.
func Calculate(input string) string {
	// Удаляем пробелы
	input = strings.TrimSpace(input)
	// Разделяем строку на части
	parts := strings.Fields(input)

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
