package methods

import (
	"fmt"
	"math/rand"
)

type Bonus struct{}

func NewBonus() Bonus{
	return Bonus{}
}

func (c Bonus) Pay(usd int) int {
	fmt.Println("Размер оплаты бонусами:", usd, "Bonus")

	return rand.Int()
}

func (c Bonus) Cancel(id int) {
	fmt.Println("Отмена оплаты бонусами! ID:", id)
}