package main

import "fmt"

func intSliceHandler(intSlice []int, newValue int) []int {
	return append(intSlice, newValue)
}

func main() {
	var intSlice []int

	fmt.Printf("\nmain - intSlice:%v\n", intSlice)

	intSlice = intSliceHandler(intSlice, 5)

	fmt.Printf("\nmain - intSlice after intSliceHandler():%v\n\n", intSlice)
}
