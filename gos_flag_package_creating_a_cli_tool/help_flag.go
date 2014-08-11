package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var help bool

	var helpText = `
    Usage:

        --help               displays usage information

	--version            displays utility version

        --posixly-correct    boolean value; if set, --help and --version
                             function as described. By default, this 
                             utility parses them as strings. 

        --limit              specifies number of times arguments will be
                             be displayed.
`

	flag.BoolVar(&help, "help", false, helpText)

	flag.Parse()

	if help {
		fmt.Printf("%s\n", flag.Lookup("help").Usage)
		os.Exit(0)
	}
}
