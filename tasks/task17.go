package tasks

import "fmt"

/*
17.	Реализовать бинарный поиск встроенными методами языка
*/

func Task17() {
	// отсортированный срез
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	findNumber := 38
	result, searchCount := binarySearch(searchField, findNumber)
	if result < 0 {
		fmt.Printf("Your number %d is not in the list\n", findNumber)
	} else {
		fmt.Printf("Your number was found in position %d after %d steps of binary search\n\n", result, searchCount)
	}
}

// определяем именованные результаты result, searchCount (принимают значения по умолчанию для типов)
func binarySearch(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2

	switch {
	case len(a) == 0:
		// если после деления среза пополам его длина равна нулю, то искомый элемент не найден в срезе
		result = -1 // возвращаем несуществующий индекс -1
	case a[mid] > search: // если искомый элемент меньше, чем средний (mid), ищем в срезе, слева от среднего (mid) элемента
		result, searchCount = binarySearch(a[:mid], search)
	case a[mid] < search: // аналогично, если больше - ищем справа
		result, searchCount = binarySearch(a[mid+1:], search)
		// если в результате рекурсивных вычислений вернулся положительный индекс (т.е. элемент был найден)
		// то к текущей позиции mid (средний элемент) добавляем позиции среднего элемента в срезах из рекурсивных вызовов
		if result >= 0 {
			result += mid + 1
			/*
				1. [2 5 8 12 16 23 38 56 72 91], ищем 91
				                 ^(центр), его позиция mid=5. Ищем справа от mid (т.к. 91 > 23)      	| result = 3 + 5 + 1* = 9
				2. [38 56 72 91]
				           ^(центр), его позиция mid=2. Ищем справа от mid (т.к. 91 > 72)             	| result = 0 + 2 + 1* = 3
				3. [91]
				     ^(центр), его позиция mid=0. По ветке default result = mid. Возвращаемся обратно 	| result = 0

					* дополнительно прибавляем 1 к result, т.к. передавая правый или левый срез в рекурсивную функцию
					  индекс этого среза начинается с нуля.
					  добавляем эту единицу только когда рассматриваем правый срез. Когда рассматриваем левый срез,
					  значит искомый элемент находится в срезе, индексы которого соответствуют исходному срезу на
					  первой итерации. На второй и последующей - уже не соответствуют, но в соответствии с алгоритмом,
					  прибавлять 1 не нужно
			*/
		}
	default: // a[mid] == search
		result = mid // нашли искомый индекс
	}
	searchCount++
	return result, searchCount
}
