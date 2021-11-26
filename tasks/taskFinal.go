package tasks

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Дан массив из 10 строк, посчитать в каждой строке количество уникальных символов с помощью горутин
*/

func TaskFinal() {
	const (
		arrSize    = 10
		goroutines = 3
		quotaLimit = 2
	)

	array := [arrSize]string{"apple", "binary", "cat", "dog", "bug", "fox", "application", "goroutine", "dataset", "pencil"}

	wg := &sync.WaitGroup{}
	// канал для задач с квотой quotaLimit
	workerInput := make(chan string, quotaLimit)

	// создаем горутины
	for i := 0; i < goroutines; i++ {
		go uniqSymbols(i, workerInput, wg)
	}

	// воркеры сами разберутся - задача идет в канал.
	// рандомная горутина принимает к работе
	for _, num := range array {
		// создаем счетчик задач для решения, чтобы дождаться корректного завершения горутин
		wg.Add(1)
		workerInput <- num
	}

	// обязательно закрыть канал (пул воркеров) - иначе не дождемся окончания
	// работы воркеров. Это может привести к дедлоку или утечки памяти
	close(workerInput)

	// ожидание завершения работы горутин
	wg.Wait()
}

func uniqSymbols(workerName int, in <-chan string, wg *sync.WaitGroup) {
	for {
		// str - значение из канала, а more - bool переменная, равная false, если канал закрыт
		// горутина завершается, если канал in закрыт
		str, more := <-in
		if more {
			var unique int
			set := make(map[rune]bool)
			// цикл с range - не по байтам, а по символам
			for _, c := range str {
				//fmt.Println(set)
				if !set[c] {
					unique++
					set[c] = true
				}
			}

			fmt.Printf("Worker #%d: in string '%s' %d unique symbols\n", workerName, str, unique)
			/*
				уменьшаем счетчик задач для решения.
				defer - использовать обязательно (!), т.к. если оставить wg.Done() без defer,
				то после выполнения всех задач сразу выполнится wg.Wait(), который не станет ждать завершения горутин,
				т.к. счетчик выполненных задач станет равным нулю. Будет выведен результат. А сообщение о завершении
				горутин - нет.
				т.е. горутины не успеют завершиться до окончания работы программы. Завершиться - выполнить ветку else с
				явным возвратом return (и выводом на экран уведомления о завершении работы)
			*/
			defer wg.Done()
			runtime.Gosched()
		} else {
			fmt.Printf("All jobs is done. (Worker #%d) \n", workerName)
			return
		}
	}
}
