package subtract

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strings"
)

// Subtract Удаляет все вхождения второй строки из первой и вызывает trimStringsAfter40.
func Subtract(a, b string) string {
	// Проверяем, содержит ли строка a строку b
	if !strings.Contains(a, b) {
		// Если строка a не содержит строку b, обрезаем строку a до 40 символов и возвращаем её
		return trim.TrimStringsAfter40(a)
	}
	// Если строка a содержит строку b, удаляем все вхождения строки b из строки a
	return trim.TrimStringsAfter40(strings.ReplaceAll(a, b, ""))
}
