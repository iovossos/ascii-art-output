package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

var (
	colorFlag   string   // What color are we using?
	text2color  string   // Which characters will be colored?
	outputFlag  string   // Do we have an output file to save the result? Which one?
	alignFlag   string   // Do we use a special alignment?
	reverseFlag string   // Do we need to reverse a file? ATTENTION DO NOT DO THAT YET
	PFargs      []string // The arguments we have post-flags (after we remove the flags)
	iofInput    int      // Which of the PFargs is the input
	LenPFargs   int      // How many should the PFargs be
	banner      string   // Which file to use as a banner

)

// Parse for flags and options

func parseFlagsAndOptions() bool {
	flag.StringVar(&colorFlag, "color", "", "Set the text color")
	flag.StringVar(&outputFlag, "output", "", "Output file")
	flag.StringVar(&alignFlag, "align", "left", "Text alignment (left, center, right)")
	flag.StringVar(&reverseFlag, "reverse", "", "Reverse this file")

	flag.Parse()
	PFargs = flag.Args()

	// This is where the input is
	iofInput = 0
	if colorFlag != "" {
		iofInput = 1
		text2color = PFargs[0] //If there is text to color, the input moves one position to the right
	}

	LenPFargs = iofInput + 1 //Readjust what the legit len of PFargs should be. I need this var because of the next conditional

	// The len of PFargs depends on whether or not we are using an extra argument in the end for one of the banners or not.
	// so for example , "Hello" shadow is ok but "Hello" fuckoff is not.
	// I made it so this error would caught while checking for arguments
	// because that is the only way to catch "Hello" "World" shadow as wrong. But to be more demure,
	// I will also add a specialized error for not finding the bannner
	if len(PFargs) == 1 {
		banner = "standard.txt"
		return true
	}

	if iofInput+1 == len(PFargs)-2 {
		if PFargs[iofInput+1] == "standard" || PFargs[iofInput+1] == "shadow" || PFargs[iofInput+1] == "thinkertoy" {
			banner = PFargs[iofInput+1] + ".txt"
			LenPFargs += 1
		} else {
			fmt.Printf("Error: %s is not a valid banner option.\nPlease choose between \"Standard\",\"Shadow\" or \"Thinkertoy\".", PFargs[iofInput+1])
			return false
		}
	}
	return true
}

// Check the validity of the input
func checkValidity() bool {
	if len(PFargs) != LenPFargs {
		fmt.Println("Error: Incorrent number of arguments. Read README.md for more information")
		return false
	}
	if isNotASCII(PFargs[iofInput]) {
		fmt.Println("Error: Only ASCII characters or newline symbols (\\n) are allowed.")
		return false
	}

	if _, err := os.Stat(banner); err != nil {
		fmt.Println("File does not exist!")
	}

	return true

}

// Check if the input contains non-ASCII characters
func isNotASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}
