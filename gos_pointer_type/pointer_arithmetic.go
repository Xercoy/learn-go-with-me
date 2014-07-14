package main

import (
	"fmt"
	"unsafe"
)

func main() {
	intArray := [...]int{1, 2}
	intPtr := &intArray[0]
	var addressHolder uintptr

	fmt.Printf("\nintArray:")
	for key, value := range intArray {
		fmt.Printf("\n[%d] = %d, address %p", key, value, &intArray[key])
	}

	fmt.Printf("\n\nintPtr(*int) is pointing to address %p, dereferenced value %d\n", intPtr, *intPtr)

	// Convert a pointer to an integer to an unsafe.Pointer, then to a uintptr.
	addressHolder = uintptr(unsafe.Pointer(intPtr))

	// Increment the value of the address by the number of bytes of an element which is an integer.
	addressHolder = addressHolder + unsafe.Sizeof(intArray[0])

	// Convert a uintptr to an unsafe.Pointer, then to a pointer to an integer.
	intPtr = (*int)(unsafe.Pointer(addressHolder))

	fmt.Printf("Incremented intPtr by %d bytes, it is now pointing to %p, dereferenced value %d\n\n", unsafe.Sizeof(intArray[0]), intPtr, *intPtr)
}
