package tasks

import "fmt"

/*
14.	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

type TypeDefiner interface {
}

func getType(td TypeDefiner) string {
	return fmt.Sprintf("Type is %T", td)
}

func Task14() {
	var typeDef TypeDefiner

	typeDef = 0
	fmt.Println(getType(typeDef))

	typeDef = "abc"
	fmt.Println(getType(typeDef))

	typeDef = true
	fmt.Println(getType(typeDef))

	typeDef = make(chan int)
	fmt.Println(getType(typeDef))
}
