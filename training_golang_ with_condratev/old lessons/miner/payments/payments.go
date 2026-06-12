package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct{
	paymentsInfo map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo: make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p PaymentModule) Pay(description string, usd int) int {

	id := p.paymentMethod.Pay(usd)  // проведение оплаты и получение её id

	info := PaymentInfo{
		Description: description,
		Usd:       	 usd,
		Cancelled:   false,
	}

	p.paymentsInfo[id] = info    // Сохранение инфы об оплате

	return id
}

func (p PaymentModule) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.paymentsInfo[id] = info
}

func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

func (p PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo{
		tempMap[k] = v
	}
	return tempMap
}