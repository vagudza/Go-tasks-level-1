package tasks

import (
	"fmt"
	"sync"
)

/*
Task 41. Реализовать рефактор кода так, чтобы все числа из первой горутины были выведены батчами во второй горутине

func main() {
	produce := func() <-chan int {
		r := make(chan int)
		go func() {
			for i := 0; i < 43; i++ {
				r <- i
			}
		}()
		return r
	}

	process := func(in <-chan int) {
		const part = 5
		var buf []int

		for v := range in {
			if len(buf) == part {
				go func(buf []int) {
					for _, v := range buf {
						fmt.Print(v, " ")
					}
				}()
				buf = buf[0:0]
			}
			buf = append(buf, v)
		}
	}

	ch := produce()
	process(ch)
}
*/

func Task41() {
	wg := sync.WaitGroup{}

	produce := func() <-chan int {
		r := make(chan int)
		go func() {
			for i := 0; i < 43; i++ {
				r <- i
			}
			close(r) // point 1. Close channel - it is need for correct closing for v := range in  .... below
		}()
		return r
	}

	process := func(in <-chan int) {
		const part = 5
		var buf, tempBuf []int

		for v := range in {
			if len(buf) == part {
				// var tempBuf []int		// point 2. copy(nilSlice, source) - nilSlice will be empty
				// var tempBuf []int = []int{}	// some mistake - len(tempBuf) == 0, no space for copy

				tempBuf = make([]int, part)
				copy(tempBuf, buf) // point 3. It is need to copy buf data, because below buf = buf[0:0]

				wg.Add(1) // point 4. Add wg.Add()
				go func(buf []int) {
					defer wg.Done()

					for _, v := range buf {
						fmt.Print(v, " ")
					}
				}(tempBuf) // point 5. Throw tempBuf as param

				buf = buf[0:0]
			}
			buf = append(buf, v)
		}

		// optional: output tail
		if len(buf) > 0 {
			for _, v := range buf {
				fmt.Print(v, " ")
			}
		}
	}

	ch := produce()
	process(ch)

	wg.Wait()
}
