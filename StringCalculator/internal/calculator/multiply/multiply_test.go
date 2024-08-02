package multiply

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMultiply_ValidInput(t *testing.T) {
	result := Multiply("A", "3")
	assert.Equal(t, "AAA", result)
}

func TestMultiply_InvalidNumber(t *testing.T) {
	require.PanicsWithValue(t, "Некорректное число для умножения", func() {
		Multiply("A", "abc")
	})
}

func TestMultiply_NumberTooLow(t *testing.T) {
	require.PanicsWithValue(t, "Некорректное число для умножения", func() {
		Multiply("A", "0")
	})
}

func TestMultiply_NumberTooHigh(t *testing.T) {
	require.PanicsWithValue(t, "Некорректное число для умножения", func() {
		Multiply("A", "11")
	})
}
