package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var mtx = sync.Mutex{}
var money = atomic.Int64{}  // usd
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request){

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Fail to read HTTP body:", err)
		return
	}

	paymentAmount, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		fmt.Println("error convert value:", err)
		return
	}

	mtx.Lock()
	if money.Load() - int64(paymentAmount) >= 0 {
		time.Sleep(3*time.Second)
		money.Add(int64(-paymentAmount))
		fmt.Println("The payment was successful!", money.Load())
	}else {
		fmt.Println("Not enough money to pay")
	}
	mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request){
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Fail to read http request body:", err)
		return
	}

	saveAmount, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		fmt.Println("Fail to convert http body to integer:", err)
		return
	}

	mtx.Lock()
	if int64(saveAmount) <= money.Load() && int64(saveAmount) > 0 {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("Money:", money.Load(), "Bank:", bank.Load())
	}else if int64(saveAmount) < 0{
		fmt.Println("You need promblem??? You CHEATER!!!")
	}else{
		fmt.Println("Not enough money to put in bank")
	}
	mtx.Unlock()

}

func main() {

	money.Add(150)

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}

}