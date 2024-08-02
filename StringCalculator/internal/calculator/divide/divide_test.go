package divide

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDivide_ValidInput(t *testing.T) {
	result := Divide("AAAAA", "5")
	assert.Equal(t, "A", result)
}

func TestDivide_InvalidNumber(t *testing.T) {
	require.PanicsWithValue(t, "Некорректное число для деления", func() {
		Divide("AAAAA", "abc")
	})
}

func TestDivide_NumberTooLow(t *testing.T) {
	require.PanicsWithValue(t, "Некорректное число для деления", func() {
		Divide("AAAAA", "0")
	})
}

func TestDivide_NumberTooHigh(t *testing.T) {
	require.PanicsWithValue(t, "Некорректное число для деления", func() {
		Divide("AAAAA", "11")
	})
}

func TestDivide_StringTooShort(t *testing.T) {
	result := Divide("A", "2")
	assert.Equal(t, "", result)
}
