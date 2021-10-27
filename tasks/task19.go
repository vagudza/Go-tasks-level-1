package tasks

import (
	"fmt"
	"unicode/utf8"
)

/*
19.	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func Task19() {
	/*
		Для представления Юникода в Go используется тип rune, который иначе называется int32.
		Строки в Go закодированы в UTF-8. Один код может использовать 8, 16 или 32 бита памяти.
	*/
	question := "¿Cómo estás?"
	fmt.Printf("%s\n%s", question, ReverseString(question))
}

func ReverseString(s string) string {
	// функция для определения длины строки в рунах, а не байтах. Символы не из ASCII могут представлять более 1 байта
	strLen := utf8.RuneCountInString(s)

	// создаем срез рун, определяем длину в рунах
	reverseString := make([]rune, strLen)
	ind := 1
	// цикл с range - не по байтам, а по символам
	for _, c := range s {
		reverseString[strLen-ind] = c
		ind++
	}
	return string(reverseString)
}
