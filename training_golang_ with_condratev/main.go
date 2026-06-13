package main

import ( 
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

var mtx = sync.Mutex{}
var money = atomic.Int64{}  // usd
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("HTTP method: ", r.Method)

	for k, v := range r.Header{
		fmt.Println("k: ", k, "--v: ", v)
	}
	
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		msg := "Fail to read HTTP body:" + err.Error()
		fmt.Println(msg)

		_, err := w.Write([]byte(msg))
		if err != nil{
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	paymentAmount, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "error convert value:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock()
	if money.Load() - int64(paymentAmount) >= 0 {
		money.Add(int64(-paymentAmount))
		msg := "The payment was successful! " + strconv.Itoa(int(money.Load()))
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil{
			fmt.Println("fail to write HTTP response:", err)
		}
	}else {
		fmt.Println("Not enough money to pay")
	}
	mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("HTTP method: ", r.Method)

	for k, v := range r.Header{
		fmt.Println("k: ", k, "--v: ", v)
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Fail to read http request body: " + string(err.Error())
		_, err := w.Write([]byte(msg))
		if err != nil{
			fmt.Println("fail to write HTTP response: ", err)
		}

		fmt.Println(msg)
		return
	}

	saveAmount, err := strconv.Atoi(string(httpRequestBody))
	if err != nil {
		msg := "Fail to convert http body to integer: " + string(err.Error())
		_, err := w.Write([]byte(msg))
		if err != nil{
			fmt.Println("fail to write HTTp response: ", err)
		}
		fmt.Println(msg)
		return
	}

	mtx.Lock()
	if int64(saveAmount) <= money.Load() && int64(saveAmount) > 0 {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		msg := "Money: "+ strconv.Itoa(int(money.Load())) + " Bank: " + strconv.Itoa(int(bank.Load()))
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response:", err)		}
		
		fmt.Println(msg)
	}else if int64(saveAmount) < 0{
		msg := "You need promblem??? You CHEATER!!!"
		fmt.Println(msg)
		w.Write([]byte(msg))
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