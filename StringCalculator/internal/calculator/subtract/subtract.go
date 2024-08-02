package subtract

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strings"
)

func Subtract(stringOne, stringTwo string) string {

	if !strings.Contains(stringOne, stringTwo) {

		return trim.TrimStringsAfter40(stringOne)
	}

	return trim.TrimStringsAfter40(strings.ReplaceAll(stringOne, stringTwo, ""))
}
