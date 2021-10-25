package tasks

import "fmt"

/*
11.	Реализовать пересечение двух неупорядоченных множеств.
*/

func Task11() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{4, 5, 6, 7}

	// создаем импровизированное множество в виде карты
	set1 := createSet(array1)
	set2 := createSet(array2)

	resultSet := setIntersection(set1, set2)

	// вывод в виде "множества"
	prettySetOutput(resultSet)
}

// метод создает карту с уникальными ключами (множество)
func createSet(arr []int) map[int]bool {
	set := make(map[int]bool)
	for _, val := range arr {
		set[val] = true
	}
	return set
}

func setIntersection(set1, set2 map[int]bool) map[int]bool {
	set := make(map[int]bool)
	// сохраняем ключи из первого множества, которые имеются во втором множестве
	for key := range set1 {
		if set2[key] {
			set[key] = true
		}
	}
	return set
}

func prettySetOutput(set map[int]bool) {
	fmt.Print("(")
	for key := range set {
		fmt.Printf("%v,", key)
	}
	fmt.Print(")")
}
