package trim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimStringsAfter40(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is a short string.", "This is a short string."},
		{"This string is exactly forty characters long!!!", "This string is exactly forty characters ..."},
		{"Short", "Short"},
		{"This string has forty-one characters!!!", "This string has forty-one characters!!!"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, TrimStringsAfter40(test.input))
	}
}
