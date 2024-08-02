package main

import (
	"fmt"
	"github.com/Eastwesser/KataStringCalc/internal/calculator"
)

func main() {
	// Объявляем переменную input для хранения пользовательского ввода
	var input string

	// Запрос ввода от пользователя
	fmt.Print("Введите выражение: ")

	// Считываем ввод пользователя и сохраняем его в переменную input
	_, err := fmt.Scanln(&input) // Считываем ввод пользователя и сохраняем его в переменную input
	// Проверяем, произошла ли ошибка при считывании ввода
	if err != nil {
		// Если ошибка есть, вызываем панику с сообщением об ошибке
		panic("Ошибка! Введите валидное выражение!!!")
	}
	// Вызываем функцию Calculate из пакета calculator с введенным выражением
	result := calculator.Calculate(input)
	// Выводим результат на экран
	fmt.Println(result)
}
