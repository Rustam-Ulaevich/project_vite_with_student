package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank{
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Размер оплаты:", usd, "Dollars")

	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("Отмена банковской операции! ID:", id)
}