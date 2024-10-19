package main

func main() {

	// Parse the arguments for flags and options
	if !parseFlagsAndOptions() {
		return
	}
	// Check the validity of the input
	if !checkValidity() {
		return
	}
	// Load the banner (ASCII art)
	asciiMap, asciiHeight := loadBanner(banner)
	// Process the input string
	processString(PFargs[iofInput], asciiMap, asciiHeight)
}
