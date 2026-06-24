package main

import (
	"context"
	"fmt"
	"study/miner"
	"study/postman"
	"sync"
	"sync/atomic"
	"time"
)


func main() {
	var tasks atomic.Int64

	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3*time.Second)
		minerCancel()
		fmt.Println("The ---> personnel's <--- working day is legal !!!!")
	}()
	go func() {
		time.Sleep(6*time.Second)
		postmanCancel()
		fmt.Println("The ---> postman's <--- working day is legal !!!!!!!!")
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 300)
	mailTransferPoint := postman.PostmanPool(postmanContext, 300)

	initTime := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(1)               // Вариант с использованием горутин
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			tasks.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	// isCoalClosed := false
	// isMailClosed := false

	// for !isCoalClosed || !isMailClosed{
	// 	select {		
	// 	case c, ok := <- coalTransferPoint:
	// 		if !ok {
	// 			isCoalClosed = true
	// 			continue
	// 		}
	// 		tasks += c
	// 	case m, ok := <- mailTransferPoint:
	// 		if !ok {
	// 			isMailClosed = true
	// 			continue
	// 		}
	// 		mails = append(mails, m)
	// 	}
	// }
	
	fmt.Println("Number of tasks:", tasks.Load())

	mtx.Lock()
	fmt.Println("Number of emails received:", len(mails))
	mtx.Unlock()

	fmt.Println("Затраченное время:", time.Since(initTime))
}