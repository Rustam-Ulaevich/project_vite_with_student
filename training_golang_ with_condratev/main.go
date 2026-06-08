package main

import (
	"context"
	"fmt"
	"time"
)

// Parent context
func parentChild(ctx context.Context, n int) {
	for{
		select{
		case <- ctx.Done():
			fmt.Println("parentChild has ended-----!!!", n)
			return
		default:
			fmt.Println("parentChild continues", n)
		}		
		time.Sleep(200*time.Millisecond)
	}
}

// Child context
func childFunc(ctx context.Context, n int) {
	for{
		select{
		case <- ctx.Done():   // возвращает канал, который закрывается при отмене контекста
			fmt.Println("childFunc has ended-----!!!", n)
			return
		default:   // срабатывает, если ctx.Done() еще не закрыт
			fmt.Println("childFunc continues", n)// Без sleep горутина бы спамила сообщениями каждую наносекунду
		}
		time.Sleep(200*time.Millisecond)
	}
}


func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go parentChild(parentContext, 1)
	go parentChild(parentContext, 2)
	go parentChild(parentContext, 3)
	go childFunc(childContext, 1)
	go childFunc(childContext, 2)
	go childFunc(childContext, 3)

	time.Sleep(1*time.Second)
	childCancel()

	time.Sleep(1*time.Second)
	parentCancel()
	
	time.Sleep(3*time.Second)
	fmt.Println("Main the end")
}