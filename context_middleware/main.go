package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("server started!")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greeting", middleware(handler))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	traceID := r.Context().Value("traceID")
	w.Write([]byte(traceID.(string)))
}

func middleware(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "traceID", "a1s2d3f4g5")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	//defer fmt.Println("after signal interrupt")

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Hola, hoy es martes, como estas?")
	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		fmt.Printf("error is %s \n", err.Error())
	}
}
