package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
6. Реализовать все возможные способы остановки выполнения горутины
*/

func Task6() {
	ch1 := make(chan string)
	run := true

	go func(run *bool) {
		for {
			select {
			// 1 способ: дождаться закрытия канала
			case <-ch1:
				fmt.Println("Goroutine 1 done (<-ch1)!")
				return
			default:
				// 2 способ: интересный, но пошлый
				if !*run {
					fmt.Println("Goroutine 1 done (*run == false)!")
					return
				}
				// какая-то полезная работа
				fmt.Println("Goroutine 1 working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(&run)

	time.Sleep(150 * time.Millisecond)
	run = false
	time.Sleep(350 * time.Millisecond)
	// остановка горутины по закрытию канала. Через этот же канал можно получать данные из горутины, но
	// в таком случае нужно синхронизировать момент после закрытия канала и работы горутины с каналом:
	// чтобы горутина не писала в закрытый канал (panic)
	close(ch1)

	// поэтому, более оптимальный вариант завершения горутины с каналом:
	ch2 := make(chan int)
	go func() {
		for {
			// 3 способ: используем второе возвращаемое каналом значение
			// num - значение из канала, а more - bool переменная, равная false, если канал закрыт
			num, more := <-ch2
			if !more {
				fmt.Println("Goroutine 2 done!")
				return
			}
			// какая-то полезная работа
			fmt.Printf("Goroutine 2 working says %d\n", num)
		}
	}()

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)

	time.Sleep(350 * time.Millisecond)

	// другим вариантом остановки является использование таймера
	// у таймера есть канал, и в случае если мы будем использовать мультиплексор select, то мы можем поставить на
	// чтение с этого канала, и как только наступит нужное время, там появится событие, и он сработает.

	// рассмотрим пример с бесконечным циклом в горутине
	timer := time.NewTimer(500 * time.Millisecond)
	timeout := time.After(300 * time.Millisecond)
	go func() {
		for {
			select {
			case <-timer.C: // 4 способ: канал таймера time.NewTimer
				fmt.Println("Goroutine 3: timer.C timeout happened")
				return
			// У NewTimer есть короткое объявление, когда нам сразу же возвращается канал без
			// промежуточной переменной:
			// 		case <-time.After(time.Millisecond):
			//			// do something
			// Но у них есть небольшие особенности. Мы можем остановить timer от
			// выполнения (!), а time.After мы остановить не можем (!), и пока он не выполнится, даже если
			// мы завершили функцию, он не освободит ресурсы. Поэтому если вам нужно лучше контролировать
			// расход ресурсов, расход памяти, пользуйтесь простым таймером и останавливайте его, когда вам
			// это требуется.
			// В данном примере case <-time.After(time.Millisecond) - никогда не выполнится, поскольку
			// функция time.After(time.Millisecond) создается внутри бесконечного цикла с default, который
			// будет выполняться раньше, чем выполнится таймер

			case <-timeout: // 5 способ: time.After (еще аналогичные варианты: time.NewTicker(), time.Tick())
				// пока не выстрелит - не соберется сборщиком мусора
				fmt.Println("Goroutine 3: timer.After timeout happened")
				return

				// default:
				// 	// какая-то полезная работа
				// 	fmt.Println("Goroutine 3 working...")
				// 	time.Sleep(1000 * time.Millisecond)
			}
		}
	}()
	time.Sleep(650 * time.Millisecond)

	// еще пример без бесконечного цикла в горутине
	timer1 := time.NewTimer(300 * time.Millisecond)
	go func() {
		select {
		case <-timer1.C: // канал таймера
			fmt.Println("Goroutine 4: timer.C timeout happened")
			return

		case <-time.After(400 * time.Millisecond):
			// пока не выстрелит - не соберется сборщиком мусора
			fmt.Println("Goroutine 4: timer.After timeout happened")
			return

		case result := <-longSQLQuery(): // 6 способ: канал из отдельной функции
			fmt.Println("Goroutine 4: longSQLQuery() result:", result)
			// освобождает ресурс, если longSQLQuery() выполнится раньше
			if !timer1.Stop() {
				<-timer.C
			}
		}
	}()
	time.Sleep(650 * time.Millisecond)

	// 7 способ: использование пакета context (WithCancel, WithDeadline, WithTimeout)
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1)

	// запуск 3 горутин, останавливаем их по завершению работы первой из них
	for i := 0; i <= 3; i++ {
		go worker(ctx, i, result)
	}

	// Дожидаемся первого результата
	foundBy := <-result
	fmt.Println("result found by", foundBy)

	// завершаем работу горутин
	finish()
}

func longSQLQuery() chan int {
	// созаем буферизированный канал, чтобы положить в буфер значение и вернуться в основную горутину
	// иначе если использовать небуферизированный канал, то он примет значение, которое никто не читает (panic)
	out := make(chan int, 1)
	time.Sleep(150 * time.Millisecond)
	out <- 100
	return out
}

func worker(ctx context.Context, workerNum int, out chan<- int) {
	// имитация некоторой работы:
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println("worker", workerNum, "sleep", waitTime)

	select {
	// БЕЗОПАСНО ИСПОЛЬЗОВАТЬ между неск. горутинами (ctx.Done)

	// либо завершаем горутину по сигналу от finish()
	case <-ctx.Done():
		return

	// либо по времени
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}
