package tasks

/*
Task 42. В приведенном коде ниже написать функцию merge так, чтобы она вернула отсортированный по возрастанию
слайс. На входе в функцию merge - два канала с отсотированными данными

// want -1, 0, 1, 2, 3, 3, 5, 9, 10, 10, 12, 15
func mergeChannels(ch1 <-chan int64, ch2 <-chan int64) []int64 {
	// TODO
}

func initChannel(data []int64) chan int64 {
	ch := make(chan int64, len(data))
	for _, d := range data {
		ch <- d
	}
	close(ch)

	return ch
}

func Task42() {
	ch1 := initChannel([]int64{1, 2, 3, 10, 12, 15})
	ch2 := initChannel([]int64{-1, 0, 3, 5, 9, 10})

	fmt.Println(mergeChannels(ch1, ch2))
}

*/

import (
	"fmt"
)

// want -1, 0, 1, 2, 3, 3, 5, 9, 10, 10, 12, 15
func mergeChannels(ch1 <-chan int64, ch2 <-chan int64) []int64 {
	var result []int64

	var1, ok1 := <-ch1
	var2, ok2 := <-ch2

	for {
		if ok1 && ok2 {
			if var1 < var2 {
				result = append(result, var1)
				var1, ok1 = <-ch1
			} else {
				result = append(result, var2)
				var2, ok2 = <-ch2
			}
		} else {
			if !ok1 {
				result = append(result, var2)
				var2, ok2 = <-ch2
			}

			if !ok2 {
				result = append(result, var1)
				var1, ok1 = <-ch1
			}

			if !ok1 && !ok2 {
				break
			}
		}
	}

	return result
}

func initChannel(data []int64) chan int64 {
	ch := make(chan int64, len(data))
	for _, d := range data {
		ch <- d
	}
	close(ch)

	return ch
}

func Task42() {
	ch1 := initChannel([]int64{1, 2, 3, 10, 12, 15})
	ch2 := initChannel([]int64{-1, 0, 3, 5, 9, 10})

	fmt.Println(mergeChannels(ch1, ch2))
}
