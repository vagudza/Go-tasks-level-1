package tasks

import (
	"fmt"
	"sync"
)

/*
18.	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	sum int
	mu  *sync.Mutex
}

func (c *Counter) Add() {
	// блокируем любой доступ к c.sum на время инкремента
	c.mu.Lock()
	c.sum++
	c.mu.Unlock()
}

func (c *Counter) GetSum() int {
	return c.sum
}

// создаем "конструктор"
func newCounter() *Counter {
	return &Counter{mu: &sync.Mutex{}}
}

func Task18() {
	c := newCounter()
	wg := &sync.WaitGroup{}

	// запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerNum int) {
			for j := 0; j < 1000; j++ {
				c.Add()
			}
			wg.Done()
			fmt.Printf("Goroutine #%d is done\n", workerNum)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Result is %d", c.GetSum())
}
