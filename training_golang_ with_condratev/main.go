package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Println("Действие №", n, "началось...")
	time.Sleep(1*time.Second)
	fmt.Println("Действие №", n, "закончилось...")

	transferPoint <- 10                                // Отправка числа 10 в канал
	fmt.Println("Действие №", n, "Число передано!!!")
}


func main() {
	coal := 0

	transferPoint := make(chan int)  // создание канала

	initTime := time.Now()

	go mine(transferPoint, 1)    // асинхронный вызов (в отдельной горутине)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <- transferPoint   //  получение значения из канала
	coal += <- transferPoint
	coal += <- transferPoint

	fmt.Println("Накоплено", coal, "единиц")
	fmt.Println("Прошло времени:", time.Since(initTime))
}