package tasks

import (
	"fmt"
	"time"
)

/*
25.	Реализовать собственную функцию sleep.
*/

func Task25() {
	fmt.Println("Start sleeping...")
	Sleep(1)
	fmt.Println("Wake up!")
}

func Sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}
