package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	var name string

	stdinFileInfo, _ := os.Stdin.Stat()

	// Assign content piped to stdin if possible, assign random name if not.
	if (stdinFileInfo.Mode() & os.ModeNamedPipe) != 0 {

		// Content is being piped, assign it to name
		stdinContent, _ := ioutil.ReadAll(os.Stdin)

		name = string(stdinContent)
	} else {
		// No content was piped, choose a random salutation.
		defaultNames := []string{"user", "mysterious person",
			"misinformed user who should be piping their name through"}

		// Randomization stuff, ensures seed is different on every run.
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(defaultNames))

		name = defaultNames[randomIndex]
	}

	fmt.Printf("Hello, %s!", name)
}
