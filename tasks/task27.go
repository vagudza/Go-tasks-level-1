package tasks

/*
27. В приведенной ниже программе на языке Go дописать функцию merge так, чтобы она
выдавала в стандартный вывод упорядоченные значения,  считанные из каналов -
аргументов за исключением тех, что повторяются. Предполагается, что в каналы
подаются строго упорядоченные по возрастанию значения без повторений (см функцию dump).
Следует рассматривать вариант как конечных, так и бесконечных каналов (на которых не вызывается close).
В приведенном примере должно быть выведено  3 5 6 7 24 50


package main

import ("fmt" "math")

func dump(c chan int, a []int) {
    for _, i:= range(a) { c <- i }
    close(c)
}

func merge(c1, c2 chan int) {
    var1, ok1 := <-c1
    var2, ok2 := <-c2
    ...
}

func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    a1 := [] int {1,5,7,8,24,25,50}
    a2 := [] int {1,3,6,8,25}

    go dump(c1, a1)
    go dump(c2, a2)

    merge(c1, c2)
}
*/

import (
	"fmt"
)

func dump(c chan int, a []int) {
	for _, i := range a {
		c <- i
	}
	close(c)
}

// merge: main alg of this func - find first min value from both channel and print it. After that, printed value will removed
// from buffer. If values from both buffers is equal, remove this values
func merge(c1, c2 chan int) {
	var (
		var1, var2 int
		ok1, ok2   bool
		buf1, buf2 []int
	)

	// important: if not set ok1 and ok2, next case can be:
	// go dump(c2, a2) dumps all second slice a2, and first slice a1 at this time is not dumped yet
	// then the programm outputs all second slice, because ok1 is false by default (var ok1 bool)
	// and when second slice was printed, ok1 still be false (because ch1 was closed)
	ok1, ok2 = true, true

	for {
		select {
		case var1, ok1 = <-c1:
			if ok1 {
				buf1 = append(buf1, var1)
			}
		case var2, ok2 = <-c2:
			if ok2 {
				buf2 = append(buf2, var2)
			}
		}

		// notice: if one of channels was closed, next case is can be:
		// reads from closed channel, outs from select and runs code below
		// it adds some unuseful iterations

		buf1, buf2 = process(buf1, buf2)

		// in case of closing both channels (work is done)
		if !(ok1 || ok2) {
			// it needs to output remaining values from buf1 or buf2
			if len(buf1) > 0 {
				printSlice(buf1)
			} else {
				printSlice(buf2)
			}

			break
		}
	}
}

// process select min element from buf1 and buf2, then prints it. After that remove min elem and returns new buf1 and buf2
func process(buf1, buf2 []int) ([]int, []int) {
	if len(buf1) == 0 || len(buf2) == 0 {
		return buf1, buf2
	}

	switch {
	case buf1[0] == buf2[0]:
		return buf1[1:], buf2[1:]
	case buf1[0] > buf2[0]:
		idx := printBefore(buf2, buf1[0])
		buf2 = buf2[idx+1:]
	case buf2[0] > buf1[0]:
		idx := printBefore(buf1, buf2[0])
		buf1 = buf1[idx+1:]
	}

	return buf1, buf2
}

// printBefore prints all values from buf, which is less or equal than limit
// returns max index of buf's slice: buf[index] <= limit
func printBefore(buf []int, limit int) int {
	var i, n int
	for i, n = range buf {
		if n >= limit {
			return i - 1
		}

		fmt.Print(n, " ")
	}

	return i
}

func printSlice(s []int) {
	for _, n := range s {
		fmt.Print(n, " ")
	}
}

func Task27() {
	c1 := make(chan int)
	c2 := make(chan int)

	a1 := []int{1, 5, 7, 8, 24, 25, 50}
	a2 := []int{1, 3, 6, 8, 25}

	go dump(c1, a1)
	go dump(c2, a2)

	merge(c1, c2)
}
