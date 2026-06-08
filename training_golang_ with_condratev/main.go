package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

// var number int = 0  // 10000
// var number atomic.Int64
var slice []int

var mtx sync.Mutex  // sync.Mutex - взаимное исключение (mutual exclusion)

func increment(wg *sync.WaitGroup){
	defer wg.Done()
	
	for i := 1; i <= 1000; i++{
		// number.Add(1)
		mtx.Lock()                //блокирует мьютекс
		slice = append(slice, i)
		mtx.Unlock()             // разблокирует мьютекс. Позволяет другой ожидающей горутине войти в критическую секцию	
	}	
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(10)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)
	go increment(wg)

	wg.Wait()

	// fmt.Println("Main the end", number.Load())	
	fmt.Println("Main the end", len(slice))	

}