package main

import (
	"KataStringCalc/internal/calculator"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаем новый сканер для считывания ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Запрос ввода от пользователя
	fmt.Print("Введите выражение: ")

	// Считываем всю строку ввода
	scanner.Scan()
	input := scanner.Text()

	// Проверяем, что ввод не пустой и содержит хотя бы один пробел
	if strings.TrimSpace(input) == "" || !strings.Contains(input, " ") {
		fmt.Println("Ошибка! Введите валидное выражение!!!")
		return
	}

	// Вызываем функцию Calculate из пакета calculator с введенным выражением
	result := calculator.Calculate(input)
	// Выводим результат на экран
	fmt.Println(result)
}
