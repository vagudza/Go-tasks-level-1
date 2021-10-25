package tasks

import (
	"fmt"
)

/*
9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout
*/
const arraySize = 5

func Task9() {
	array := [arraySize]int{2, 4, 6, 8, 10}

	// устанавливаем пайплайн, который возвращает канал
	ch := gen(array)
	// далее этот канал передаем в следующую функцию расчета квадратов, которая возвращает канал
	out := sq(ch)

	// потребляем вывод
	for val := range out {
		fmt.Println(val) // вывод в stdout
	}
}

// первая функция из конвеера чисел: запись в канал чисел из массива
func gen(nums [arraySize]int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// вторая функция из конвеера чисел: запись в канал квадрата чисел
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
