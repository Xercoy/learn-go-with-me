package main

import (
	"fmt"
	"os"
)

func main() {
	for true {
		if len(os.Args) == 1 {
			fmt.Printf("y \n")
		} else {
			break;
		}
	}
}
