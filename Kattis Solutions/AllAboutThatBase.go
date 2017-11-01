package main

import (
	"strings"
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	// Repeat for input n test cases
	for i := 0; i < n; i++ {
		var s1, operation, s2, answer string
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
		// Include special case (base 1) if input only consists of the digit 1
		for j, b1 := 0, true; j < len(testCase); j++ {
			if string(testCase[j]) != "1" {
				b1 = false
			}
			if j == len(testCase)-1 && b1 {
				max = 1
			}
		}

		/* 	Convert to all possible bases and evaluate the expression.
			If true for a given base, append the base as a digit between
			1 and z for bases 1 through 35, or 0 for base 36, to a string
			containing all possible answers. Keep track of whether or not
			a solution has been found using the boolean solutionExists.
			If no solution is found, return "invalid" for the expression,
			otherwise return all the valid bases.
		 */
		solutionExists := false
		answers := ""
		for j := max; j <= 36; j++ {

			n1, n2, nAns := toInt(s1, genBase(j)), toInt(s2, genBase(j)), toInt(answer, genBase(j))
			if operation == "/" {
				if n1%n2 != 0 {
					continue
				}
			}
			if nAns == opByString(operation, n1, n2) {
				solutionExists = true
				if j != 36 {
					answers += genBase(36)[j:j+1]
				} else {
					answers += "0"
				}
			}
		}

		if !solutionExists {
			fmt.Println("invalid")
		} else {
			fmt.Println(answers)
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
