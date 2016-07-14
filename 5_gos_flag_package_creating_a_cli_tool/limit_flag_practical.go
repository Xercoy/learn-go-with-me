package main

import(
"fmt"
"flag"
"strconv"
"errors"
)

type customFlag struct {
	FlagValue int
}

func (cF *customFlag) String() string {
	return strconv.Itoa(cF.FlagValue)
}

func (cF *customFlag) Set(s string) error {
	convertedArg, _ := strconv.Atoi(s)

	if (convertedArg < 0) {
		return errors.New("Argument value for limit flag invalid. Must be greater than or equal to 0.")
	}

	cF.FlagValue = convertedArg

	return nil
}

func (cF *customFlag) Get() int {
	return cF.FlagValue
}

func main() {
	limit := customFlag{-1}

	flag.Var(&limit, "limit", "Limits the number of times arguments should be displayed. Infinity otherwise.")

	flag.Parse()

	if limit.Get() == -1 {
		for true {
			fmt.Printf("y \n")
		}
	} else {
		for i := 0; i < limit.Get(); i++ {
			fmt.Printf("y \n")
		}
	}
}