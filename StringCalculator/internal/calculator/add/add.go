package add

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
)

// Add Складывает две строки и вызывает trimStringsAfter40.
func Add(a, b string) string {
	return trim.TrimStringsAfter40(a + b)
}
