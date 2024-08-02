package subtract

// Пакет subtract содержит функции для вычитания строк

import (
	"github.com/Eastwesser/KataStringCalc/internal/calculator/trim"
	"strings"
)

/*
trim для обрезки строк
strings для работы со строками
*/

// Subtract Удаляет все вхождения второй строки из первой и вызывает trimStringsAfter40.
func Subtract(stringOne, stringTwo string) string {
	/*
		Проверяем, содержит ли строка stringOne строку stringTwo с помощью функции strings.Contains
		Если строка stringOne не содержит строку stringTwo, функция вернет false
	*/
	if !strings.Contains(stringOne, stringTwo) {
		/*
			Если строка stringOne не содержит строку stringTwo,
			обрезаем строку stringOne до 40 символов с помощью функции trim.TrimStringsAfter40
			и возвращаем обрезанную строку
		*/
		return trim.TrimStringsAfter40(stringOne)
	}

	/*
		Если строка stringOne содержит строку stringTwo,
		удаляем все вхождения строки stringTwo из строки stringOne с помощью функции strings.ReplaceAll:

		Первый аргумент - строка, в которой производится замена
		Второй аргумент - подстрока для поиска
		Третий аргумент - подстрока для замены (в данном случае пустая строка для удаления)
		Результат обрезается до 40 символов с помощью trim.TrimStringsAfter40 и возвращается
	*/
	return trim.TrimStringsAfter40(strings.ReplaceAll(stringOne, stringTwo, ""))
}
