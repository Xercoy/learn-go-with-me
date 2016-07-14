package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	versionUsageText = "0.1"
	helpUsageText    = `
  Usage: yes [FLAGS] [ARGUMENTS]

      --help        displays usage information

      --version     displays utility version

      --limit       specifies number of times 
                    arguments will be
                    displayed. Value must be
                    0 or greater.

  Note: flags must be specified before arguments.
`

	limitUsageText = "0 or greater integer value, determines times output is displayed"
)

type customFlag struct {
	Value int
}

func (cF *customFlag) Set(s string) error {
	cF.Value, _ = strconv.Atoi(s)
	return nil
}

func (cF *customFlag) Get() int {
	return cF.Value
}

func (cF *customFlag) String() string {
	return strconv.Itoa(cF.Value)
}

func (cF *customFlag) IsSet() bool {
	if cF.Value == -1 {
		return false
	}

	return true
}

func main() {
	versionFlagPtr := flag.Bool("version", false, versionUsageText)

	var helpFlagValue bool
	flag.BoolVar(&helpFlagValue, "help", false, helpUsageText)

	//-1 is the default value, see the IsSet method.
	limitFlagCustom := customFlag{-1}
	flag.Var(&limitFlagCustom, "limit", limitUsageText)

	var outputString string

	flag.Parse()

	// Remember, this needs to be executed after flag.Parse()
	argCount := len(flag.Args())

	// Append arguments together to make output string
	for _, argument := range flag.Args() {
		outputString += argument + " "
	}

	switch {

	// Version Flag
	case *versionFlagPtr:
		{
			fmt.Printf("Version: %s\n", flag.Lookup("version").Usage)
			break
		}

	//Help Flag
	case helpFlagValue:
		{
			fmt.Printf("%s\n", flag.Lookup("help").Usage)
			break
		}

	//Limit Flag
	case limitFlagCustom.IsSet() == true:
		{
			//Limit flag set w.o any arguments, print y limited times
			if argCount == 0 {
				for i := 0; i < limitFlagCustom.Get(); i++ {
					fmt.Printf("y \n")
				}

				// Limit flag set with arguments, print outputString limited times
			} else {
				for i := 0; i < limitFlagCustom.Get(); i++ {
					fmt.Printf("%s\n", outputString)
				}
			}

			break
		}

	//Default, no flags specified.
	case true:
		{
			// Default without any arguments
			if argCount == 0 {
				for true {
					fmt.Printf("y \n")
				}

				// Default with arguments
			} else {
				for true {
					fmt.Printf("%s\n", outputString)
				}
			}
		}
	}
}
