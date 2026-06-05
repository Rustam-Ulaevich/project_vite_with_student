package main


type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}


//   ---Slice---

type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

func (m PaymentModuleWithSlice) addInfo(info PaymentInfo) {

}

//  ---Map---

type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}


func main() {

}