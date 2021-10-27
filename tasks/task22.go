package tasks

import (
	"fmt"
	"math/big"
)

/*
22.	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func Task22() {
	/*
		Пакет big предоставляет три типа данных:

		    big.Int для крупных целых чисел, когда 18 квинтиллионов недостаточно;
		    big.Float для вещественных чисел с плавающей запятой производной точности;
		    big.Rat для дробей вроде 1/3.

		еще один вариант решения - использовать срезы байт, длина которых соответствует количеству разрядов в больших числах
	*/
	s1 := "900000000"
	s2 := "100000000"

	// создаем числа из строк
	a, _ := new(big.Float).SetString(s1)
	if a == nil {
		fmt.Printf("Error get float from string\n")
		return
	}
	b, _ := new(big.Float).SetString(s2)
	if b == nil {
		fmt.Printf("Error get float from string\n")
		return
	}

	sum := new(big.Float)
	sum.Add(a, b)

	sub := new(big.Float)
	sub.Sub(a, b)

	div := new(big.Float)
	div.Quo(a, b)

	mult := new(big.Float)
	mult.Mul(a, b)

	fmt.Println("Sum is ", sum)
	fmt.Println("Sub is ", sub)
	fmt.Println("Div is ", div)
	fmt.Println("Mult is ", mult)
}
