package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) { // wg нужен для синхронизации горутин с main
	defer wg.Done()  // выполнится ПЕРЕД выходом из функции

	for i := 1; i <= 3; i++{
		fmt.Println("Я отнёс", text, "в", i, "раз")
		time.Sleep(200*time.Millisecond)
	}	
}


func main() {

	wg := sync.WaitGroup{}   // Создаём WaitGroup (счётчик синхронизации)

	wg.Add(1)                  // увеличивает счётчик WaitGroup на 1
	go postman(&wg, "газету")

	wg.Add(1)
	go postman(&wg, "журнал")

	wg.Add(1)
	go postman(&wg, "письмо")

	wg.Wait()  //БЛОКИРУЕТ выполнение main. Ожидает, пока счётчик WaitGroup не станет равен 0

	fmt.Println("Main the end")

}