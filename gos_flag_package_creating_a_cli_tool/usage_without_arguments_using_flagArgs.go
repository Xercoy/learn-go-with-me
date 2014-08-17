package main

import (
	"fmt"
"flag"
)

func main() {
	flag.Parse()

	for true {
		if len(flag.Args()) == 0 {
			fmt.Printf("y \n")
		} else {
			break;
		}
	}
}
