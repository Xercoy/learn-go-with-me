package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	for true {
		if len(flag.Args()) == 0 {
			fmt.Printf("y \n")
		} else {
			break
		}
	}
}
