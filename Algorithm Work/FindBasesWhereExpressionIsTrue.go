package main

import (
	"strings"
	"fmt"
	"math"
)

func main() {

	fmt.Print("\n\tWelcome to Ciarán Ó hAoláin's program to generate a list of bases where a given expression is true, written in Go!\n\n\tStart by typing the number of test cases, and hitting return!\n\n\t")

	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Printf("\n\n\t%d test cases? Awesome! Now type each test case in the format '123 * 4be = 5fs' and hit return to print the result!\n\n", n)

	// Repeat for input n test cases
	for i := 0; i < n; i++ {
		var s1, operation, s2, answer string
		fmt.Print("\t")
		fmt.Scanf("%s %s %s = %s\n", &s1, &operation, &s2, &answer)
		testCase := s1 + s2 + answer

		// Find the base at which to start
		max := 0
		for j := 0; j < len(testCase); j++ {
			high := strings.LastIndex(genBase(36), string(testCase[j])) + 1
			if high > max {
				max = high
			}
		}

		// Consider special case (base 1) where input only consists of the digit 1
		for j, b1 := 0, true; j < len(testCase); j++ {
			if string(testCase[j]) != "1" {
				b1 = false
			}
			if j == len(testCase)-1 && b1 {
				max = 1
			}
		}

		/* 	Convert to all possible bases and evaluate the expression.
			If true for a given base, append the base as a digit to a string
			containing all possible answers. Keep track of whether or not
			a solution has been found using the boolean solutionExists.
			If no solution is found, return "invalid" for the expression,
			otherwise return all the valid bases.
		 */
		solutionExists := false
		var answers []int
		for j := max; j <= 36; j++ {

			// Convert all values in expression to integers for current base
			n1, n2, nAns := toInt(s1, genBase(j)), toInt(s2, genBase(j)), toInt(answer, genBase(j))

			// Consider scenarios where integer division will lead to a false positive
			if operation == "/" {
				if n1%n2 != 0 {
					continue
				}
			}

			// Check if expression is true
			if nAns == opByString(operation, n1, n2) {
				solutionExists = true
				answers = append(answers, j)
			}
		}

		// Print results for expression
		if !solutionExists {
			fmt.Println("\tinvalid\n")
		} else {
			fmt.Print("\tExpression is true in base(s):\t")
			for j := 0; j < len(answers)-1; j++ {
				fmt.Printf("%d, ", answers[j])
			}
			fmt.Println(answers[len(answers)-1], "\n")
		}

	}

}

func toInt(input, source string) (outputInt int) {

	baseSource := len(source)

	if baseSource == 1 {
		return len(input)
	}

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

func opByString(opString string, a, b int) int {

	switch opString {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}

	return -1

}
