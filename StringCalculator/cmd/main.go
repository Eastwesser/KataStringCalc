package main

import (
	"fmt"
	"github.com/Eastwesser/KataStringCalc/internal/calculator"
)

func main() {
	var input string
	fmt.Print("Введите выражение: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("Ошибка ввода")
	}

	result := calculator.Calculate(input)
	fmt.Println(result)
}
