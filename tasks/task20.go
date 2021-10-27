package tasks

import (
	"fmt"
	"regexp"
)

/*
20.	Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

*/

func Task20() {
	s := "snow        dog sun 1 2 3 ¿Cómo estás?"
	fmt.Println(s)
	fmt.Println(ReverseSentense(s))
}

func ReverseSentense(s string) string {
	var reverse string
	// находим все слова, разделенные непробельными символами
	reFindWords, _ := regexp.Compile(`(\S)+`)
	words := reFindWords.FindAllString(s, -1)

	// находим все пробельные символы [\r\n\t\f\v ]
	reSpaces, _ := regexp.Compile(`(\s)+`)
	spaces := reSpaces.FindAllString(s, -1)

	// инвертируем, с учетом пробельных символов
	wordsLen := len(words)
	for i := 0; i < wordsLen; i++ {
		reverse += words[wordsLen-i-1]
		if i < wordsLen-1 {
			reverse += spaces[wordsLen-i-2]
		}
	}

	return reverse
}
