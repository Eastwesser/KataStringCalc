package add

// Объявляем пакет add, который содержит функции для сложения строк

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
)

// trim для обрезки строк

// Add Складывает две строки и вызывает trimStringsAfter40.
func Add(stringOne, stringTwo string) string {
	/*
		Складываем строки stringOne и stringTwo с помощью оператора '+'
		Результат представляет собой новую строку, состоящую из конкатенации stringOne и stringTwo
	*/
	concatenatedString := stringOne + stringTwo

	/*
		Передаем полученную строку в функцию trim.TrimStringsAfter40
		Эта функция обрежет строку до 40 символов, если она длиннее
		И вернет результат
	*/
	return trim.TrimStringsAfter40(concatenatedString)
}
