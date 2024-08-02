package add

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
)

func Add(stringOne, stringTwo string) string {
	concatenatedString := stringOne + stringTwo

	return trim.TrimStringsAfter40(concatenatedString)
}
