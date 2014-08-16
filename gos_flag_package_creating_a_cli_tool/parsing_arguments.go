package main

import (
	"flag"
	"fmt"
)

func main() {

	for _, argument := range flag.Args() {
		fmt.Printf("\n%s", argument)
	}
}
