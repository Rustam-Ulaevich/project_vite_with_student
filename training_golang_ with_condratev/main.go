package main

import (
	"context"
	"fmt"
	"study/miner"
	"study/postman"
	"time"
)


func main() {
	var tasks int
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

	coalTransferPoint := miner.MinerPool(minerContext, 2)
	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	isCoalClosed := false
	isMailClosed := false

	for !isCoalClosed || !isMailClosed{
		select {		
		case c, ok := <- coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}
			tasks += c
		case m, ok := <- mailTransferPoint:
			if !ok {
				isMailClosed = true
				continue
			}
			mails = append(mails, m)
		}
	}
	
	fmt.Println("Number of tasks:", tasks)
	fmt.Println("Number of emails received:", tasks)

}