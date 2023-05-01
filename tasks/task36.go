package tasks

import "fmt"

/*
Task 36. Найти максимальное и следующее за максимальным в массиве. N > 1
*/

func Task36() {
	maxIdx := 0
	secondMaxIdx := 1

	a := []int{10, 11, 7, 5, 11, 6, 7, 12, 11}
	if a[0] < a[1] {
		maxIdx, secondMaxIdx = secondMaxIdx, maxIdx
	}

	for i := 2; i < len(a); i++ {
		if a[i] > a[maxIdx] {
			secondMaxIdx = maxIdx
			maxIdx = i
		} else {
			if a[i] > a[secondMaxIdx] {
				secondMaxIdx = i
			}
		}
	}

	fmt.Printf("Task 36. max=%d; second max=%d \n", a[maxIdx], a[secondMaxIdx])
}
