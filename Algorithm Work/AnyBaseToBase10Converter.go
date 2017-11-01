package main

import (
	"strings"
	"fmt"
	"math"
	"os"
)

func main() {

	fmt.Printf("\n\tWelcome to Ciarán Ó hAoláin's Base Conversion program, written in Go!\n\n\tEnter number in any base between base 1 and base 36 in the format 'number base', \n\tfollowed by return, to convert to base 10!\n\n\t")
	var input string
	var base int
	fmt.Scanf("%s %d\n", &input, &base)
	input = strings.ToLower(input)

	// Error checking
	if base > 36 || base < 1 {
		fmt.Print("\n\n\t	Error:\tIllegal base provided\n\n")
		os.Exit(1)
	}
	for i := 0; i < len(input); i++ {
		if !(strings.Contains(genBase(base), string(input[i]))) {
			fmt.Print("\n\n\tError:\tInput contains characters not present in provided base\n\n")
			os.Exit(1)
		}
	}

	// Print answer
	if base != 1 {
		fmt.Printf("\n\tAnswer: %d\n\n\t", toInt(input, genBase(base)))
	} else {
		fmt.Printf("\n\tAnswer: %d\n\n\t", len(input))
	}

}

func toInt(input, source string) (outputInt int) {

	baseSource := len(source)

	for j, indexInSource := len(input)-1, -1; j >= 0; j-- {
		indexInSource = strings.Index(source, string(input[j]))
		outputInt += indexInSource * intPower(baseSource, len(input)-j-1)
	}

	return

}

func intPower(a, b int) int {

	return int(math.Pow(float64(a), float64(b)))

}

func genBase(base int) string {

	mostBases := "0123456789abcdefghijklmnopqrstuvwxyz"
	if base == 1 {
		return "1"
	} else {
		return mostBases[0:base]
	}

}
