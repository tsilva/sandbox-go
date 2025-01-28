package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the range of numbers to multiply
	const numElements = 10000000

	// Start timer
	start := time.Now()

	// Perform the multiplication
	result := 1.0 // Using float64 to avoid overflow for demonstration
	for i := 1; i <= numElements; i++ {
		result *= float64(i)
		if result == 0.0 {
			fmt.Println("Result too large and became zero due to floating point precision.")
			break
		}
	}

	// Stop timer
	elapsed := time.Since(start)

	// Output the result and time elapsed
	fmt.Printf("Time elapsed: %s\n", elapsed)
	fmt.Printf("Final result (or partial): %g\n", result)
}
