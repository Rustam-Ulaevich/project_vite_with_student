package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal{
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Размер оплаты через PayPal:", usd, "USD")

	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("Отмена транзакции PayPal! ID:", id)
}