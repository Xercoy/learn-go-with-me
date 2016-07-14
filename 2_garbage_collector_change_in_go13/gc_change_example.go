package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	var intPtr1, intPtr2 *int
	var addressHolder uintptr

	intPtr1 = new(int)
	*intPtr1 = 100

	intPtr2 = new(int)
	*intPtr2 = 999

	fmt.Printf("\nintPtr1 = %p, dereferenced value = %d\n", intPtr1, *intPtr1)
	fmt.Printf("intPtr2 = %p, dereferenced value = %d\n", intPtr2, *intPtr2)

	addressHolder = uintptr(unsafe.Pointer(intPtr1))
	fmt.Printf("\naddressHolder(uintptr) is assigned the integer value of the address %p.\n", unsafe.Pointer(addressHolder))

	intPtr1 = intPtr2
	fmt.Printf("\nintPtr1 is assigned the value of inPtr2, value = %p, dereferenced value = %d\n", intPtr1, *intPtr1)

	fmt.Printf("\nCalling runtime.GC()\n")
	runtime.GC()

	fmt.Printf("\naddressHolder(uintptr), value = %p, dereferenced value = %d\n\n", unsafe.Pointer(addressHolder), *(*int)(unsafe.Pointer(addressHolder)))
}
