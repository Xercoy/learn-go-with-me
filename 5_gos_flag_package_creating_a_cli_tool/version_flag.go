package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var version *bool
	version = flag.Bool("version", false, "1.0")

	flag.Parse()

	if *version {
		fmt.Printf("Version %s\n", flag.Lookup("version").Usage)
		os.Exit(0)
	}
}
