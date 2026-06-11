package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(
	ctx context.Context, 
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int) {
		defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Personnel number:", n, "The working day is over.")
			return
		default :
			fmt.Println("Personnel number:", n, "started worked.")
			time.Sleep(1*time.Second)
			fmt.Println("Personnel number:", n, "completed the task:", power)

			transferPoint <- power
			fmt.Println("Personnel number:", n, "passed the task:", power)
		}		
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int{
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i:=1 ; i<=minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg,  coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}