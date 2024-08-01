package main

import (
	"fmt"
	"github.com/Eastwesser/KataStringCalc/internal/calculator"
)

func main() {
	var input string

	fmt.Print("Введите выражение: ") // Запрос ввода

	_, err := fmt.Scanln(&input) // Считывание в переменную input
	if err != nil {
		panic("Ошибка ввода")
	}

	result := calculator.Calculate(input) // вызов функции Calculate из пакета calculator
	fmt.Println(result)                   // Вывод результата
}
