package main

import "fmt"

type Foo struct {
	intSlice []int
}

func Fill(f Foo, newValue int) {
	fmt.Printf("\nFill - Executing append(f.intSlice, newValue)")
	f.intSlice = append(f.intSlice, newValue)
	
	fmt.Printf("\nFill - f.intSlice:%v", f.intSlice)
}

func main() {
	var FooVariable Foo
	
	fmt.Printf("\nmain - FooVariable.intSlice:%v", FooVariable.intSlice)

	fmt.Printf("\nmain - Executing Fill(FooVariable, 12)")
	Fill(FooVariable, 12)

	fmt.Printf("\nmain - FooVariable.intSlice:%v\n\n", FooVariable.intSlice)
}