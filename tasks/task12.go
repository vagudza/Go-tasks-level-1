package tasks

import (
	"fmt"
	"strings"
)

/*
12.	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func Task12() {
	ds1 := Dataset1{data: []string{"cat", "cat", "dog", "cat", "tree"}}
	ds2 := Dataset2{data: "cat,cat,dog,cat,tree"}

	prettySetStringsOutput(createSetFromAny(ds1))
	prettySetStringsOutput(createSetFromAny(ds2))
}

type Set interface {
	toSlice() []string
}

// Некоторая структура, удовлетворяющая интерфейсу Set
type Dataset1 struct {
	data []string
}

// реализация интерфейса Set
func (d Dataset1) toSlice() []string {
	return d.data
}

// Еще одна некоторая структура, удовлетворяющая интерфейсу Set
type Dataset2 struct {
	data string
}

// реализация интерфейса Set
func (d Dataset2) toSlice() []string {
	return strings.Split(d.data, ",")
}

// метод создает карту с уникальными ключами (множество)
func createSetFromAny(dataArray Set) map[string]bool {
	set := make(map[string]bool)

	// используя реализацию интерфейса, итерируемся и получаем значение
	for _, val := range dataArray.toSlice() {
		set[val] = true
	}
	return set
}

// вывод в stdout содержимого множества
func prettySetStringsOutput(set map[string]bool) {
	fmt.Print("(")
	for key := range set {
		fmt.Printf("%v,", key)
	}
	fmt.Print(")")
}
