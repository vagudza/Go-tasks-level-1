package tasks

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
Task 33. Последовательно идущие единицы

Требуется найти в бинарном векторе самую длинную последовательность единиц и вывести её длину.
Желательно получить решение, работающее за линейное время и при этом проходящее по входному массиву только один раз.
Формат ввода
Первая строка входного файла содержит одно число n, n ≤ 10000. Каждая из следующих n строк содержит ровно одно число — очередной элемент массива.
Формат вывода
Выходной файл должен содержать единственное число — длину самой длинной последовательности единиц во входном массиве.
*/

func Task33() {
	fileIn, err := os.Open("tasks/inputs/input33.txt")
	if err != nil {
		log.Fatal("open err=", err)
	}

	var num, currentLen, maxLen int
	var skipFirstIteration bool
	scanner := bufio.NewScanner(fileIn)
	for scanner.Scan() {
		if !skipFirstIteration {
			skipFirstIteration = true
			continue
		}

		num, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("convert err=", err)
		}

		if num == 1 {
			currentLen++

			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner err=", err)
	}

	fileOut, err := os.OpenFile("tasks/outputs/output33.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open file err=", err)
	}

	dw := bufio.NewWriter(fileOut)
	_, err = dw.WriteString(strconv.Itoa(maxLen))
	if err != nil {
		log.Fatal("write string err=", err)
	}
	dw.Flush()
}
