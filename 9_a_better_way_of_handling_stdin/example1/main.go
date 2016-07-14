package main

import (
	"log"
	"os"
)

func main() {
	stdinFileInfo, _ := os.Stdin.Stat()
	log.Printf("os.Stdin File Mode  : %s", stdinFileInfo.Mode().String())
}
