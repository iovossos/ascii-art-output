package main

import "os"

func main() {
	// Check the validity of the input
	if !checkValidity() {
		return
	}
	// Load the banner (ASCII art)
	asciiMap, asciiHeight := loadBanner("standard.txt")
	// Process the input string
	processString(os.Args[1], asciiMap, asciiHeight)
}
