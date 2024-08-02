package divide

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strconv"
)

// Divide Делит первую строку на указанное количество частей и вызывает trimStringsAfter40.
func Divide(a, b string) string {
	n, err := strconv.Atoi(b)
	if err != nil {
		panic("Некорректное число для деления")
	}

	if n < 1 || n > 10 {
		panic("Некорректное число для деления")
	}

	if len(a) < n {
		return ""
	}

	return trim.TrimStringsAfter40(a[:len(a)/n])
}
