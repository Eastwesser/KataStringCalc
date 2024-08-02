package divide

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strconv"
)

func Divide(stringOne, stringTwo string) string {
	n, err := strconv.Atoi(stringTwo)

	if err != nil {
		panic("Некорректное число для деления")
	}

	if n < 1 || n > 10 {
		panic("Некорректное число для деления")
	}

	if len(stringOne) < n {
		return ""
	}

	return trim.TrimStringsAfter40(stringOne[:len(stringOne)/n])
}
