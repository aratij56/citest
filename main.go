// main.go
package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello World!")
	result := Add(3, 4)
	fmt.Println("Sum:", result)
}

// Add is a simple function that adds two numbers
func Add(a, b int) int {
	return a + b
}
