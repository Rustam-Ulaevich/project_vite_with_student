package main

import (
	"fmt"
	"net/http"
	"time"
)

func payHandler(w http.ResponseWriter, r *http.Request){
	str := "Pay order"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("An error occurred while recording the HTTP response:", err.Error())
	}else{
		fmt.Println("Pay was processed correctly")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request){
	time.Sleep(3*time.Second)
	str := "Cansel pay"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("An error occurred while recording the HTTP response:", err.Error())
	}else{
		fmt.Println("Cansel pay was processed correctly")
	}
}

func handler(w http.ResponseWriter, r *http.Request){
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("An error occurred while recording the HTTP response:", err.Error())
	} else {
		fmt.Println("HTTP request was processed correctly")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Start HTTP server")
	err := http.ListenAndServe(":9091", nil)
	if err !=nil {
		fmt.Println("An error has occurred:", err.Error())
	}
	fmt.Println("Finish HTTP server")
}
