package add

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		stringOne, stringTwo, expected string
	}{
		{"Hello", "World", "HelloWorld"},
		{"Foo", "Bar", "FooBar"},
		{"", "Test", "Test"},
		{"Test", "", "Test"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Add(test.stringOne, test.stringTwo))
	}
}
