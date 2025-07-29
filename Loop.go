package main

import (
	"fmt"
)

func main() {
	fruits := []string{"Apple", "Banana", "Cherry"}

	for i, fruit := range fruits {
		fmt.Printf("Fruit %d: %s\n", i+1, fruit)
	}
}
