package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set a seed for the current time in milliseconds.
	// Ensures random values on every run.
	rand.Seed(time.Now().UnixNano())

	randomByteSlice := make([]byte, 100)

	// Start timer
	startTime := time.Now()

	rand.Read(randomByteSlice)

	// Retrieve total run time
	endTime := time.Now().Sub(startTime)

	fmt.Printf("%s - created random slice %v",
		endTime.String(), randomByteSlice)
}
