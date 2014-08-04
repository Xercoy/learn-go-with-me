package main

import "fmt"

func intSliceHandler(intSlice []int, newValue int) {
	intSlice = make([]int, 1)
	intSlice[0] = newValue

	fmt.Printf("\nintSliceHandler - intSlice:%v\n", intSlice)
}

func main() {
	var intSlice []int

	fmt.Printf("\nmain - intSlice:%v\n", intSlice)

	intSliceHandler(intSlice, 5)

	fmt.Printf("\nmain - intSlice after intSliceHandler():%v\n\n", intSlice)
}
