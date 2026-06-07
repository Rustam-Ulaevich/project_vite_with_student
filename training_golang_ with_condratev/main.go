package main

import (
	"fmt"
	"strconv"
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

	

	



	/*
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