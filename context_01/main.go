package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background() //creamos padre

	key, value := "ID", "a98d73b87"

	ctx = addValue(ctx, key, value)

	printValue(ctx, key)
}

func addValue(ctx context.Context, key string, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func printValue(ctx context.Context, key string) {
	valueContext := ctx.Value(key)
	fmt.Printf("%v: %v \n", key, valueContext)
}
