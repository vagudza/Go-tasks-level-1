package tasks

import "fmt"

/*
Task 39.

	Дана последовательность положительных числел длиной N и число X
	Найти два различных числа последовательности A и B, таких, что A+B=X
	Или вернуть 0,0 если таких чисел нет
*/
func Task39() {
	arr := []int{1, 4, 2, 4, 6, 2, 7}
	X := 8
	A, B := getAB(arr, X)
	fmt.Println(A, B)

	A, B = getABFast(arr, X)
	fmt.Println(A, B)
}

// Complexity: Time: O(N+N)--> O(N), Memory: O(N)
func getAB(arr []int, X int) (int, int) {
	hashTableArr := make(map[int]struct{}, len(arr))
	for i := range arr {
		hashTableArr[arr[i]] = struct{}{}
	}

	for i := range arr {
		need := X - arr[i]
		if _, ok := hashTableArr[need]; ok && need != arr[i] {
			return arr[i], need
		}
	}

	return 0, 0
}

// []int{1, 4, 2, 4, 6, 2, 7}
// Complexity: Time: O(N), Memory: O(N)
func getABFast(arr []int, X int) (int, int) {
	hashTableArr := make(map[int]struct{}, len(arr))
	for i := range arr {
		need := X - arr[i]
		if _, ok := hashTableArr[need]; ok && need != arr[i] {
			return arr[i], need
		}

		hashTableArr[arr[i]] = struct{}{}
	}

	return 0, 0
}
