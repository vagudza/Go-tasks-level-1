package tasks

import (
	"fmt"
	"os"
)

/*
8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func Task8() {
	// функция принимает число в десятичной СС, i-ая позиция бита (начиная с нуля), и значение i-го бита
	a, err := setBitByPosition(15, 2, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Result is %d (%b)", a, a)
}

func setBitByPosition(num int64, pos uint8, bit uint8) (int64, error) {
	// сдвигаем единицу на pos позиций для работы с i-ой позицией
	shift := int64(1) << pos

	switch bit {
	case 0:
		// | поразрядная дизъюнкция, Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
		// ^ поразрядное исключающее ИЛИ. Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
		res := (num | shift) ^ shift
		fmt.Printf("[%d (%b) | %d (%b)] ^ %d (%b) = %d (%b)\n", num, num, shift, shift, shift, shift, res, res)
		return res, nil
	case 1:
		res := num | shift
		fmt.Printf("%d (%b) | %d (%b) = %d (%b)\n", num, num, shift, shift, res, res)
		return res, nil
	default:
		return num, fmt.Errorf("bit must be 0 or 1")
	}
}
