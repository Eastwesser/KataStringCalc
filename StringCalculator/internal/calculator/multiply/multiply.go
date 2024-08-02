package multiply

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strconv"
	"strings"
)

func Multiply(stringOne, stringTwo string) string {
	n, err := strconv.Atoi(stringTwo)

	if err != nil {
		panic("Некорректное число для умножения")
	}

	if n < 1 || n > 10 {
		panic("Некорректное число для умножения")
	}

	return trim.TrimStringsAfter40(strings.Repeat(stringOne, n))
}
