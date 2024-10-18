package main

import (
	"fmt"
	"strings"
)

// Process the input string and print the ASCII art
func processString(input string, asciiMap map[rune][]string, asciiHeight int) {
	// Replace literal "\n" with actual newlines and split into lines
	input = strings.ReplaceAll(input, `\n`, "\n")
	inputLines := strings.Split(input, "\n")
	// Process each line separately
	for _, line := range inputLines {
		if line == "" {
			fmt.Println() // Handle empty lines (newline in the ASCII art)
			continue
		}
		// Build and print the ASCII art for the line
		asciiChars := buildAsciiArt(line, asciiMap, asciiHeight)
		printAsciiArt(asciiChars, asciiHeight)
	}
}

// Build the ASCII art for a given line of input
func buildAsciiArt(line string, asciiMap map[rune][]string, asciiHeight int) [][]string {
	var asciiChars [][]string
	for _, char := range line { // Use rune to support Unicode
		if art, exists := asciiMap[char]; exists {
			asciiChars = append(asciiChars, art)
		} else {
			// Handle characters not present in the font data
			fmt.Printf("Warning: Character '%c' not found in font data.\n", char)
			asciiChars = append(asciiChars, make([]string, asciiHeight))
		}
	}
	return asciiChars
}

// Print the ASCII art for the given characters
func printAsciiArt(asciiChars [][]string, asciiHeight int) {
	for i := 0; i < asciiHeight; i++ {
		for _, charLines := range asciiChars {
			fmt.Print(charLines[i]) // Print each line of each character
			// Uncomment the next line if you want spaces between characters
			// fmt.Print(" ")
		}
		fmt.Println() // Move to the next line of the ASCII art
	}
}
