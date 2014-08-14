package main

import(
"fmt"
"flag"
"strconv"
)

type CustomFlag struct {
	FlagValue int
}

func (cF *CustomFlag) String() string {
	return strconv.Itoa(cF.FlagValue)
}

func (cF *CustomFlag) Set(s string) error {
	cF.FlagValue, _ = strconv.Atoi(s)
	
	return nil
}

func (cF *CustomFlag) Get() int {
	return cF.FlagValue
}

func main() {
	limit := CustomFlag{-1}

	flag.Var(&limit, "limit", "Limits output")

	flag.Parse()

	if limit.Get() > -1 {
		fmt.Printf("\nLimit: %d\n\n", limit.Get())
	} else {
		fmt.Printf("\nLimit flag not included.\n\n")
	}
}