package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Stat returns file info structure describing the file.
	stdinFileInfo, _ := os.Stdin.Stat()
	stdinFileMode := stdinFileInfo.Mode()

	log.Println("File Info Mode of os.Stdin : ", stdinFileMode.String())

	// Apply the masks and determine what the file mode is.
	if (stdinFileMode & os.ModeCharDevice) != 0 {
		log.Println("os.ModeCharDevice          :  ", os.ModeCharDevice)
		log.Println("Binary AND Result          :  ", (stdinFileMode & os.ModeCharDevice))

		fmt.Println("Input from Terminal")
	} else if (stdinFileMode & os.ModeNamedPipe) != 0 {
		log.Println("os.ModeNamedPipe           : ", os.ModeNamedPipe)
		log.Println("Binary AND Result          : ", (stdinFileMode & os.ModeNamedPipe))

		fmt.Println("Input from Pipe")
	}
}
