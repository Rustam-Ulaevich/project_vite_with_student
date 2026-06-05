package main

import (
	"fmt"
	"study/payments"
	"study/payments/methods"
)

func main() {
	method := methods.NewBonus()

	paymentModule := payments.NewPaymentModule(method)


	o1 := paymentModule.Pay("car", 5600)
	o2 := paymentModule.Pay("beer", 10)
	paymentModule.Pay("house", 15000)

	allInfo := paymentModule.AllInfo()

	fmt.Println(allInfo)

	paymentModule.Cancel(o1)

	allInfo = paymentModule.AllInfo()

	fmt.Println(allInfo)

	info := paymentModule.Info(o2)

	fmt.Println(info)
}