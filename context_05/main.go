package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nictes1/context-Go/context_05/animals"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)

	animalFinishedChan := make(chan bool)

	go animals.ListAnimals(ctx, animalFinishedChan)

	select {
	case <-animalFinishedChan:
		fmt.Println("Animals Finished")
	}
}
