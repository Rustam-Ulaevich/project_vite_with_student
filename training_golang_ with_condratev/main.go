package main

import (
	"errors"
	"fmt"
	// "study/payments"
	// "study/payments/methods"
)

type User struct {
	Name string
	Ballance int
}

func Pay(user *User, usd int) error{
	if user.Ballance - usd < 0 {
		return errors.New("Недостаточно средств")
	}
	user.Ballance -= usd
	return nil
}

func main() {

	user := User{
		Name: "Mike",
		Ballance: 100,
	}

	fmt.Println(user)

	err := Pay(&user, 5)

	fmt.Println(user)

	if nil != err {
		fmt.Println("Оплаты не было! Причина:", err.Error())
	} else {
		fmt.Println("Оплата произведена")
	}


	// method := methods.NewBonus()

	// paymentModule := payments.NewPaymentModule(method)


	// o1 := paymentModule.Pay("car", 5600)
	// o2 := paymentModule.Pay("beer", 10)
	// paymentModule.Pay("house", 15000)

	// allInfo := paymentModule.AllInfo()

	// fmt.Println(allInfo)

	// paymentModule.Cancel(o1)

	// allInfo = paymentModule.AllInfo()

	// fmt.Println(allInfo)

	// info := paymentModule.Info(o2)

	// fmt.Println(info)
}