package tasks

import "fmt"

/*
28. Реализовать слияние нескольких целочисленных слайсов
*/

func mergeArrays(arrs ...[]int) []int {
	var totalArrsLen int
	for _, arr := range arrs {
		totalArrsLen += len(arr)
	}

	result := make([]int, totalArrsLen)
	var lastCopiedIdx int
	for _, arr := range arrs {
		copy(result[lastCopiedIdx:], arr)
		lastCopiedIdx += len(arr)
	}

	fmt.Println("len=", len(result), "cap=", cap(result))
	return result
}

func badMergeArrays(arrs ...[]int) []int {
	var result []int
	for _, arr := range arrs {
		result = append(result, arr...)
	}

	fmt.Println("len=", len(result), "cap=", cap(result))
	return result
}

func Task28() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5}
	arr3 := []int{6, 7, 8}

	fmt.Println(mergeArrays(arr1, arr2, arr3))    // [1 2 3 4 5 6 7 8], but len= 8 cap= 8
	fmt.Println(badMergeArrays(arr1, arr2, arr3)) // [1 2 3 4 5 6 7 8], but len= 8 cap= 12 - bad solution!
}
