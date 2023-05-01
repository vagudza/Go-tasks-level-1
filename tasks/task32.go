package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Task 32. Генерация скобочных последовательностей

Дано целое число n. Требуется вывести все правильные скобочные последовательности длины 2 ⋅ n, упорядоченные лексикографически (см. https://ru.wikipedia.org/wiki/Лексикографический_порядок).
В задаче используются только круглые скобки.
Желательно получить решение, которое работает за время, пропорциональное общему количеству правильных скобочных последовательностей в ответе, и при этом использует объём памяти, пропорциональный n.
Формат ввода
Единственная строка входного файла содержит целое число n, 0 ≤ n ≤ 11
Формат вывода
Выходной файл содержит сгенерированные правильные скобочные последовательности, упорядоченные лексикографически.
*/

func Task32() {
	n := readFile32()
	tree := []string{"("}
	newTree := make([]string, 0)

	var lp, rp int
	if n == 0 {
		tree[0] = ""
	}

	for len(tree[0]) < 2*n {
		for _, node := range tree {
			lp = leftPars(node)
			rp = rightPars(node)

			if lp < n {
				newTree = append(newTree, node+"(")
			}

			if lp > rp {
				newTree = append(newTree, node+")")
			}
		}

		tree = newTree
		newTree = nil
	}

	for _, node := range tree {
		fmt.Println(node)
	}

	writeToFile32(tree)
}

func leftPars(s string) int {
	return strings.Count(s, "(")
}

func rightPars(s string) int {
	return strings.Count(s, ")")
}

func readFile32() int {
	fileIn, err := os.Open("tasks/inputs/input32.txt")
	if err != nil {
		log.Fatal("open err=", err)
	}

	scanner := bufio.NewScanner(fileIn)
	n := 0
	if scanner.Scan() {
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("input file: can not convert line to int", err)
		}
	}

	return n
}

func writeToFile32(tree []string) {
	fileOut, err := os.OpenFile("tasks/outputs/output32.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open file err=", err)
	}

	dw := bufio.NewWriter(fileOut)

	for _, node := range tree {
		_, err = dw.WriteString(node + "\n")
		if err != nil {
			log.Fatal("write string err=", err)
		}
	}

	dw.Flush()
}
