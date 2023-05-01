package tasks

import "fmt"

/*
Task 34. Дана последовательность слов, вывести самые короткие слова через пробел
*/
func Task34() {
	s := []string{"monster", "cat", "room", "flat", "me", "you", "test", "dump", "if", "done"}

	shortIdx := make([]int, 0)
	shortLen := -1
	for i, val := range s {
		if shortLen == -1 || len(val) < shortLen {
			shortLen = len(val)
			shortIdx = shortIdx[0:0]
		}

		if len(val) == shortLen {
			shortIdx = append(shortIdx, i)
		}
	}

	for _, idx := range shortIdx {
		fmt.Println(s[idx])
	}
	fmt.Println()
}
