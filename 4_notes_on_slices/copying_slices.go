package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3}

	secondIntSlice := intSlice
	fmt.Printf("\nintSlice:%v\n", intSlice)

	secondIntSlice[0] = 9
	fmt.Printf("\nsecondIntSlice[0] = 9.\n")
	fmt.Printf("\nsecondIntSlice:%v\n", intSlice)

	fmt.Printf("\nintSlice:%v\n", intSlice)

	secondIntSlice = secondIntSlice[0:1]
	fmt.Printf("\nsliced secondIntSlice to [0,1].\n")
	fmt.Printf("\nsecondIntSlice:%v\n", secondIntSlice)

	fmt.Printf("\nintSlice - len:%v cap:%v\n", len(intSlice), cap(intSlice))
	fmt.Printf("\nsecondIntSlice - len:%v cap:%v\n\n", len(secondIntSlice), cap(secondIntSlice))
}
