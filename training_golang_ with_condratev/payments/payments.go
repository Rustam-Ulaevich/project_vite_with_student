package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct{
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) PaymentModule {
	return PaymentModule{
		paymentMethod: paymentMethod,
	}
}

// func (p PaymentModule) Pay(description string, usd int) int {
// 	id := p.paymentMethod.Pay(usd)

// 	info := PaymentInfo{
// 		Description: description,
// 		Usd:       	 usd,
// 		Cancelled:   false,
// 	}
// }

func (p PaymentModule) Cancel() {}

func (p PaymentModule) Info() {}

func (p PaymentModule) AllInfo() {}