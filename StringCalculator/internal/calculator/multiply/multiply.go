package multiply

// Пакет multiply содержит функции для умножения строк

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strconv"
	"strings"
)

/*
trim для обрезки строк
strconv для преобразования строк в числа
strings для работы со строками
*/

// Multiply Повторяет первую строку указанное количество раз и вызывает trimStringsAfter40.
func Multiply(stringOne, stringTwo string) string {
	// Преобразуем строку b в целое число n
	n, err := strconv.Atoi(stringTwo)

	// Проверяем, возникла ли ошибка при преобразовании
	if err != nil {
		// Если ошибка есть, вызываем панику с сообщением о некорректном числе для умножения
		panic("Некорректное число для умножения")
	}

	// Проверяем, находится ли n в диапазоне от 1 до 10
	if n < 1 || n > 10 {
		// Если n вне допустимого диапазона, вызываем панику с сообщением о некорректном числе для умножения
		panic("Некорректное число для умножения")
	}

	// Повторяем строку stringOne n раз, обрезаем результат до 40 символов и возвращаем
	return trim.TrimStringsAfter40(strings.Repeat(stringOne, n))
}
