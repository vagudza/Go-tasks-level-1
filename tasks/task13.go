package tasks

import "fmt"

/*
13.	Поменять местами два числа без создания временной переменной.
*/

func Task13() {
	a := 1
	b := 2
	fmt.Printf("Task13. a=%d, b=%d\n", a, b)

	a, b = b, a
	fmt.Printf("Task13. a=%d, b=%d\n", a, b)

	a, b = swap(a, b)
	fmt.Printf("Task13. a=%d, b=%d\n", a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}
