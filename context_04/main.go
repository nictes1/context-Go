package main

import (
	"context"
	"time"

	"github.com/nictes1/context-Go/context_04/message"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*4)

	message.PrintMessage(ctx, time.Second*5, "Hola, ¿Como estan?")
}
