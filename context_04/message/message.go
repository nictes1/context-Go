package message

import (
	"context"
	"fmt"
	"time"
)

func PrintMessage(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t): //cuando pase el tiempo recibe el channel y lo muesta
		fmt.Println("Message: " + msg)
	case <-ctx.Done(): //cuando finalice el contexto
		fmt.Println("Context finalizado.")
	}
}
