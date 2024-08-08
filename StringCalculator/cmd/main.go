package main

import (
	"KataStringCalc/internal/calculator"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите выражение: ")

	scanner.Scan()
	input := scanner.Text()

	if strings.TrimSpace(input) == "" || !strings.Contains(input, " ") {
		fmt.Println("Ошибка! Введите валидное выражение!!!")
		return
	}

	result := calculator.Calculate(input)
	fmt.Println(result)
}
