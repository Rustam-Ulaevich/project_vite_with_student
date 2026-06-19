package main

import (
	"fmt"
	"net/http"
)

//localhost:9091/default?v=x&t=s

func hanler(w http.ResponseWriter, r *http.Request){
	vParam := r.URL.Query().Get("v")
	tParam := r.URL.Query().Get("t")

	fmt.Println("v:", vParam)
	fmt.Println("t:", tParam)
}

func main() {
	http.HandleFunc("/default", hanler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("failed to run http server:", err)
	}

}