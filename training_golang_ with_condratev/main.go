package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
// func mine(transferPoint chan int, n int) {
// 	fmt.Println("Действие №", n, "началось...")
// 	time.Sleep(1*time.Second)
// 	fmt.Println("Действие №", n, "закончилось...")

// 	transferPoint <- 10                                // Отправка числа 10 в канал
// 	fmt.Println("Действие №", n, "Число передано!!!")
// }
*/


func main() {

	// ch := make(chan int)  // create new chanel

	// // ---------closed chanel-----------

	// close(ch)             //  closed chanel

	// //close(ch)            // !!!!!! panic: close of closed channel

	// v1 := <-ch            //read closed chanel
	// v2 := <-ch
	// v3 := <-ch

	// fmt.Println(v1, v2, v3)

	// v4, ok1 := <-ch        //read closed chanel
	// v5, ok2 := <-ch
	// v6, ok3 := <-ch

	// fmt.Println(v4, v5, v6, ok1, ok2, ok3)

	// ch <- 10               //   !!!!!!!!!!!panic: send on closed channel


	transferPoint := make(chan int)

	go func(){
		iterations := 3 + rand.Intn(4)  // random 3 + 0..1..2..3
		fmt.Println("iterations:", iterations)

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(300*time.Millisecond)
		}

		close(transferPoint)		
	}()

	coal := 0

	for{
		v, ok := <- transferPoint
		if !ok {
			fmt.Println("Все действия канала завершены")
			break
		}

		coal += v

		// fmt.Println(coal)
	}

	// fmt.Println("Итого:", coal)



	//----------nil chanel--------------

	var chNil chan string                             // nil chanel
	//var chNil chan string = make(chan string)     // initialization is need

	go func ()  {
		chNil <- "Random string"
	}()

	v := <-chNil

	fmt.Println(v)   // !!!!!!!!!!!!!!fatal error: all goroutines are asleep - deadlock!

	/*

	strChanel := make(chan string)
	intChanel := make(chan int)
	
	go func() {
		i := 0
		for{
			intChanel <- i
			i++
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		i := 0
		for{
			strChanel <- "str" + strconv.Itoa(i)
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// time.Sleep(5000 * time.Millisecond)

	for{
		select{
			case number := <- intChanel:
				fmt.Println("integer:", number)
			case string := <- strChanel:
				fmt.Println("string:", string)
			// default:
			// 	fmt.Println("Никакой канал не готов")
		}
		
	}
	
	coal := 0

	transferPoint := make(chan int)  // создание канала

	initTime := time.Now()

	go mine(transferPoint, 1)    // асинхронный вызов (в отдельной горутине)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	go func(){                  //анонимная ф-ая в горутинах
		for {
			fmt.Println("Анонимная функция")
			time.Sleep(100*time.Millisecond)
		}

	}()

	coal += <- transferPoint   //  получение значения из канала
	coal += <- transferPoint
	coal += <- transferPoint

	fmt.Println("Накоплено", coal, "единиц")
	fmt.Println("Прошло времени:", time.Since(initTime))
	*/
}