package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/saludar", saludar)
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message := "Hola, como estas?"

	var err error
	defer fmt.Println("after signal interrupt")

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "%s \n", message)

	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		fmt.Printf("error is %s \n", err.Error())
	}
}
