package tasks

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
)

/*
Task 31. Удаление дубликатов

Дан упорядоченный по неубыванию массив целых 32-разрядных чисел. Требуется удалить из него все повторения.
Желательно получить решение, которое не считывает входной файл целиком в память, т.е., использует лишь константный объем памяти в процессе работы.
Input format
Первая строка входного файла содержит единственное число n, n ≤ 1000000.
На следующих n строк расположены числа — элементы массива, по одному на строку. Числа отсортированы по неубыванию.
Output format
Выходной файл должен содержать следующие в порядке возрастания уникальные элементы входного массива.
*/

func Task31() {
	fileIn, err := os.Open("tasks/inputs/input31.txt")
	if err != nil {
		log.Fatal("open err=", err)
	}

	var (
		wg                 sync.WaitGroup
		currentNum         int
		lastNum            *int
		skipFirstIteration bool
	)
	ch := make(chan int)
	scanner := bufio.NewScanner(fileIn)

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeToFile31(ch)
	}()

	for scanner.Scan() {
		// first iteration contains number of strings
		if !skipFirstIteration {
			skipFirstIteration = true
			continue
		}

		currentNum, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("convert err=", err)
		}

		if lastNum == nil {
			lastNum = new(int)
			*lastNum = currentNum
			ch <- currentNum
			continue
		}

		if *lastNum == currentNum {
			continue
		}
		*lastNum = currentNum
		ch <- *lastNum
	}

	close(ch)
	wg.Wait()
}

func writeToFile31(ch chan int) {
	fileOut, err := os.OpenFile("tasks/outputs/output31.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open file err=", err)
	}

	dw := bufio.NewWriter(fileOut)

	for num := range ch {
		_, err = dw.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			log.Fatal("write string err=", err)
		}
	}

	dw.Flush()
}
