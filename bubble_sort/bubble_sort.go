package main

import "fmt"

func Sort(slice []int) {
	for end := (len(slice) - 1); end >= 0; end-- {
		for iteratingSlice := slice; &iteratingSlice[0] != &slice[end]; iteratingSlice = iteratingSlice[1:] {
			if iteratingSlice[0] > iteratingSlice[1] {
				elementHolder := iteratingSlice[0]
				iteratingSlice[0] = iteratingSlice[1]
				iteratingSlice[1] = elementHolder
			}
		}
	}
}

func main() {
	numberSlice := []int{1, 4, 6, 8, 7, 9, 3, 5, 2}

	fmt.Printf("\nMain, unsorted integer array: %v\n", numberSlice)

	Sort(numberSlice)

	fmt.Printf("\nMain, sorted integer array: %v\n\n", numberSlice)
}
