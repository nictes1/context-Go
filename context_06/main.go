package main

import (
	"github.com/nictes1/context-Go/context_06/server"
)

func main() {
	srv := server.New()
	srv.ListenAndServe()
}
