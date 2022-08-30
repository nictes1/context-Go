package server

import (
	"fmt"
	"net/http"
	"time"
)

func New() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Println("Contexto Cancelado")
	case <-time.After(2 * time.Second):
		w.Write([]byte("Hola, ¿como estan? \n"))
	}
}
