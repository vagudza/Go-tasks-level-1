package tasks

import (
	"fmt"
	"math"
	"wb-l1-tasks/point"
)

/*
24.	Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором
*/

// "конструктор" возвращает указатель на структуру
func NewPoint(x, y float64) *point.Point {
	p := &point.Point{}
	p.Set(x, y)
	return p
}

func Task24() {
	a := NewPoint(1, 2)
	b := NewPoint(1, 6)

	dist := math.Sqrt(math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2))
	fmt.Println(dist)
}
