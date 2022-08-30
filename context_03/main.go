package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//context 1º
	ctx := context.Background() //context raiz. NO se puede cancelar. Vacio

	//context 2º
	ctxCancel, _ := context.WithCancel(ctx) // Nos devuelve (ctx context.Context, cancel context.CancelFunc)
	//variable de tipo context.CancelFunc

	//context 3º
	ctxTimeOut, cancelTimeOut := context.WithTimeout(ctxCancel, time.Second*1)
	defer cancelTimeOut()

	select {
	case <-ctxCancel.Done():
		fmt.Println("Cancel context")
	case <-ctxTimeOut.Done(): //.Done() --> nos indica si dicho contexto termino
		fmt.Println("Time out context")

	}
}
