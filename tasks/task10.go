package tasks

import (
	"fmt"
	"math"
)

/*
10.	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc
*/

func Task10() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// создаем карту для хранения групп
	groups := make(map[float64][]float64)

	for _, temp := range temperatures {
		// определяем ключ карты. Например, если temp=-25.3 --> key == -20
		key := math.Trunc(temp/10.0) * 10.0
		groups[key] = append(groups[key], temp)
	}

	for key, value := range groups {
		fmt.Printf("%.0f:%v\n", key, value)
	}
}
