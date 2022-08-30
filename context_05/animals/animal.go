package animals

import (
	"context"
	"fmt"
	"time"
)

type Animal struct {
	Name string
	Age  int
}

var animals []Animal = []Animal{
	{Name: "Hormiga", Age: 2},
	{Name: "Canguro", Age: 11},
	{Name: "Tortuga", Age: 44},
	{Name: "Araña", Age: 24},
	{Name: "Elefante", Age: 62},
	{Name: "Perro", Age: 89},
	{Name: "Gato", Age: 600},
}

func ListAnimals(ctx context.Context, chn chan<- bool) {
	for _, animal := range animals {
		contextIsDone := false

		select {
		case <-ctx.Done():
			contextIsDone = true
		case <-time.After(time.Second):
			fmt.Println(animal)
		}

		if contextIsDone {
			fmt.Println("Context Finalizado")
			chn <- true
			break
		}

	}

	chn <- true
}
