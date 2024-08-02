package subtract

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strings"
)

// Subtract Удаляет все вхождения второй строки из первой и вызывает trimStringsAfter40.
func Subtract(a, b string) string {
	if !strings.Contains(a, b) {
		return trim.TrimStringsAfter40(a)
	}
	return trim.TrimStringsAfter40(strings.ReplaceAll(a, b, ""))
}
