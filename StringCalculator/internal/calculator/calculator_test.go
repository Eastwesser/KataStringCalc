package calculator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePanics(t *testing.T) {
	assert.Panics(t, func() { Calculate("invalid input") }, "Должна произойти паника при некорректном выражении")
	assert.Panics(t, func() { Calculate("\"hello\" ? \"world\"") }, "Должна произойти паника при недопустимой операции")
	assert.Panics(t, func() { Calculate("\"hello\" * invalid") }, "Должна произойти паника при некорректном числе для умножения")
	assert.Panics(t, func() { Calculate("\"hello\" / invalid") }, "Должна произойти паника при некорректном числе для деления")
	assert.Panics(t, func() { Calculate("\"hello\" / 0") }, "Должна произойти паника при делении на ноль")
	assert.Panics(t, func() { Calculate("\"hello\" / 11") }, "Должна произойти паника при делении на слишком большое число")
}

func TestCalculateNoPanic(t *testing.T) {
	assert.NotPanics(t, func() { Calculate("\"hello\" + \"world\"") }, "Паника не должна происходить при корректном выражении")
	assert.NotPanics(t, func() { Calculate("\"hello\" - \"l\"") }, "Паника не должна происходить при корректном выражении")
	assert.NotPanics(t, func() { Calculate("\"hello\" * 2") }, "Паника не должна происходить при корректном выражении")
	assert.NotPanics(t, func() { Calculate("\"hellohello\" / 2") }, "Паника не должна происходить при корректном выражении")
}

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}
