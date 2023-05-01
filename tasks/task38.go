package tasks

import (
	"fmt"
	"strconv"
)

/*
Task 38.
RLE alg:

	Дана строка (возможно пустая),состоящая из букв A-Z
	Нужно написать функцию RLE, которая сворачивает строку "XXXYYZ" --> "X3Y2Z"
	Если символ встречается один раз, он остается без изменений
	Если символ повторяется более одного раза,к нему добавляется количество повторений
*/
func Task38() {
	s := "XXXYYZRRRRRRRRRRRRRRR"
	fmt.Println(RLE(s))
}

func RLE(s string) string {
	if len(s) == 0 {
		return ""
	}

	lastChar := s[0]
	cnt := 1
	var ans string
	for i := 1; i < len(s); i++ {
		if s[i] == lastChar {
			cnt++
		} else {
			ans = writeAnswerPart(ans, lastChar, cnt)
			cnt = 1
			lastChar = s[i]
		}
	}

	return writeAnswerPart(ans, lastChar, cnt)
}

func writeAnswerPart(ans string, lastChar uint8, cnt int) string {
	if cnt == 1 {
		ans += string(lastChar)
	} else {
		ans += string(lastChar) + strconv.Itoa(cnt)
	}

	return ans
}
