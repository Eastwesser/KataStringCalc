package subtract

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtract(t *testing.T) {
	tests := []struct {
		stringOne, stringTwo, expected string
	}{
		{"Hello World", " World", "Hello"},
		{"Hello Hello", "Hello", " "},
		{"Hello", "World", "Hello"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Subtract(test.stringOne, test.stringTwo))
	}
}
