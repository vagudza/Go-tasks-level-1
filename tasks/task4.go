package tasks

import (
	"fmt"
	"time"
)

/*
4. Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из
канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func Task4() {
	var N int
	fmt.Println("Введите количество горутин:")
	fmt.Scanf("%d\n", &N)

	// можем использовать небуферизированный канал, т.к. "синхронно" пишем и читаем
	workerInput := make(chan int)

	// создаем горутины
	for i := 0; i < N; i++ {
		go worker4(i, workerInput)
	}

	// постоянная запись данных в канал из главного потока
	for {
		workerInput <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func worker4(workerName int, in <-chan int) {
	for {
		num := <-in
		fmt.Printf("Goroutine #%d: value: %ds\n", workerName, num)
	}
}
