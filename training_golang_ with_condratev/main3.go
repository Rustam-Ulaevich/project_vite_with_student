// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )


// type PaymentInfo struct {
// 	ID          int
// 	Description string
// 	Usd         int
// 	Cancelled   bool
// }

// //   ---Slice---

// type PaymentModuleWithSlice struct {
// 	s []PaymentInfo
// }

// func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
// 	m.s = append(m.s, info)
// }

// func (m *PaymentModuleWithSlice) GetInfo(id int) PaymentInfo {
// 	for _, info := range m.s {
// 		if info.ID == id {
// 			return info
// 		}
// 	}
// 	return PaymentInfo{}
// }

// //  ---Map---

// type PaymentModuleWithMap struct {
// 	m map[int]PaymentInfo
// }

// func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
// 	m.m[info.ID] = info
// }

// func (m *PaymentModuleWithMap) GetInfo(id int) PaymentInfo {
// 	info, ok := m.m[id]
// 	if !ok {
// 		return PaymentInfo{}
// 	}
// 	return info
// }

// func main() {

// 	pSlice := PaymentModuleWithSlice{}
// 	pMap := PaymentModuleWithMap{
// 		m: make(map[int]PaymentInfo),
// 	}

// 	iterations := 100000


// 	//  -------------AddInfo--------------------

// 	before := time.Now()

// 	for i := 0; i < iterations; i++ {
// 		info := PaymentInfo{
// 			ID: i,
// 			Description: strconv.Itoa(i),
// 		}
// 		pSlice.AddInfo(info)
// 	}

// 	fmt.Println("Slice add:", time.Since(before))

// 	before = time.Now()

// 	for i := 0; i < iterations; i++ {
// 		info := PaymentInfo{
// 			ID: i,
// 			Description: strconv.Itoa(i),
// 		}
// 		pMap.AddInfo(info)
// 	}

// 	fmt.Println("Map add:", time.Since(before))


// 	// ----------------MapInfo---------------------

// 	before = time.Now()

// 	for i := 0; i < iterations; i++ {
// 		info := pSlice.GetInfo(i)
// 		_ = info
// 	}
// 	fmt.Println("Slice get:", time.Since(before))

// 	before = time.Now()

// 	for i := 0; i < iterations; i++ {
// 		info := pMap.GetInfo(i)
// 		_ = info
// 	}
// 	fmt.Println("Map get:", time.Since(before))


// }