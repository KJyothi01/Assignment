package main

import (
	"fmt"
	"math/rand"
)

// Main function
func main() {

	// Using Intn() function
	ran_1 := rand.Intn(9)
	ran_2 := rand.Intn(99)
	ran_3 := rand.Intn(999)

	// Displaying the result
	fmt.Println("Random Number 1: ", ran_1)
	fmt.Println("Random Number 2: ", ran_2)
	fmt.Println("Random Number 3: ", ran_3)
}