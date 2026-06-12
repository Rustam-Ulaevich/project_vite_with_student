package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	n int,
	mail string,
) {

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Postman number:", n, "The working day is over.")
			return
		default:
			fmt.Println("Postman number:", n, "started worked.")
			time.Sleep(1 * time.Second)
			fmt.Println("Postman number:", n, "completed the task:", mail)

			transferPoint <- mail
			fmt.Println("Postman number:", n, "delivered the email:", mail)
		}
	}
	
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string{
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i:=1 ; i<=postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string{
	ptm := map[int]string{
		1: "mail",
		2: "journal",
		3: "newspaper",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "postcard"
	}

	return mail
}