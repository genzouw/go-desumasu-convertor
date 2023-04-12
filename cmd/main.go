package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/genzouw/go-desumasu-convertor/pkg/desumasu"
)

func main() {
	j, n, r := parseCommandLine()

	// Use bufio to read all from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var stdinText string
	for scanner.Scan() {
		stdinText += scanner.Text() + "\n"
	}

	fmt.Print(desumasu.Convert(stdinText, j, n, r))
}

// Command Line Options --------------------------------
//    : Convert to Keitai (default)
// -j : Convert to Jotai
// -n : Check "ね" at the end of the sentence
// -N : Remove "ね" at the end of the sentence
// ----------------------------------------------------

// Parse command line parameters
func parseCommandLine() (jotai bool, checkNe bool, removeNe bool) {
	jotai = false
	checkNe = true
	removeNe = true

	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "-j":
			jotai = true
		case "-n":
			removeNe = false
		case "-N":
			checkNe = false
		}
	}
	return jotai, checkNe, removeNe
}
