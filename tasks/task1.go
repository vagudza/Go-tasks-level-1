package tasks

import "fmt"

/*
1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).

Встраивание методов упрощает процесс использования методов.
*/

type Human struct {
	FirstName string
	LastName  string
	secret    string
}

func (h *Human) MyName() string {
	return fmt.Sprintf("%s %s", h.FirstName, h.LastName)
}

type Action struct {
	ActionName string
	// композиция структур
	Human
}

func Task1() {
	action := Action{
		ActionName: "humor",
		Human:      Human{FirstName: "Mr", LastName: "Bean"},
	}

	// в рамках этого пакета Private поле доступно, в других - недоступно.
	action.secret = "**************************"

	// Вызов встроенного в структуру Action метода из структуры Human
	fmt.Printf("Task1. %s doing %s \n", action.MyName(), action.ActionName)
}
