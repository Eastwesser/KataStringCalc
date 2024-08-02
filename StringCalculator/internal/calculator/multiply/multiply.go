package multiply

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strconv"
	"strings"
)

// Multiply Повторяет первую строку указанное количество раз и вызывает trimStringsAfter40.
func Multiply(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil {
		panic("Некорректное число для умножения")
	}

	if n < 1 || n > 10 {
		panic("Некорректное число для умножения")
	}

	return trim.TrimStringsAfter40(strings.Repeat(a, n))
}
