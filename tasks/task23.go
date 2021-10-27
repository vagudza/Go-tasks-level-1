package tasks

import "fmt"

/*
23.	Удалить i-ый элемент из слайса.
*/

// с сохранением исходного среза
func Remove(s []int, i int) ([]int, error) {
	if i >= 0 && i < len(s) {
		// создаем новый срез
		ret := make([]int, 0)
		// копируем до i-го элемента
		ret = append(ret, s[:i]...)
		// копируем после i-го элемента и возвращаем
		return append(ret, s[i+1:]...), nil
	} else {
		return []int{}, fmt.Errorf("invalid index")
	}
}

// Медленная версия (сохраняет порядок)
func RemoveSlow(s []int, i int) ([]int, error) {
	if i >= 0 && i < len(s) {
		// Выполнить сдвиг a[i+1:] влево на один индекс.
		copy(s[i:], s[i+1:])
		// Удалить последний элемент (записать нулевое значение)
		s[len(s)-1] = 0
		// Усечь срез
		s = s[:len(s)-1]
		return s, nil
	} else {
		return []int{}, fmt.Errorf("invalid index")
	}
}

// быстрый способ, без сохранения порядка. Изменяет исходный срез
func RemoveFast(s []int, i int) ([]int, error) {
	if i >= 0 && i < len(s) {
		// Копировать последний элемент в индекс i
		s[i] = s[len(s)-1]
		// Удалить последний элемент (записать нулевое значение)
		s[len(s)-1] = 0
		// Усечь срез
		s = s[:len(s)-1]
		return s, nil
	} else {
		return []int{}, fmt.Errorf("invalid index")
	}
}

func Task23() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 5
	fmt.Println("                  Slice s:", s)
	newSlice, err := Remove(s, i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New slice by Remove(s, %d): %v   <--answer\n", i, newSlice)
	fmt.Printf("                  Slice s: %v\n\n", s)

	i = 3
	fmt.Println("                     Slice s:", s)
	newSlice1, err := RemoveSlow(s, i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New slice by RemoveSlow(s, %d): %v   <--answer\n", i, newSlice1)
	fmt.Printf("                      Slice s: %v\n\n", s)

	i = 1
	fmt.Println("                      Slice s:", s)
	newSlice2, err := RemoveFast(s, i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New slice by RemoveFast(s, %d): %v   <--answer\n", i, newSlice2)
	fmt.Printf("                      Slice s: %v\n\n", s)
}
