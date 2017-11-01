package main

import (
	"fmt"
	"strings"
	"math"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		var input, source, target, output string
		var sourceAsInt int
		fmt.Scanf("%s %s %s\n", &input, &source, &target)

		baseSource := len(source)
		baseTarget := len(target)

		for j := len(input) - 1; j >= 0; j-- {

			indexInSource := strings.Index(source, string(input[j]))
			sourceAsInt += indexInSource * intPower(baseSource, len(input)-j-1)

		}

		for remainder := sourceAsInt; true; {

			output += string(target[remainder%baseTarget])
			remainder = (remainder - remainder%baseTarget) / baseTarget
			if remainder == 0 {
				break
			}

		}

		fmt.Printf("Case #%d: %s\n", i+1, reverseString(output))

	}

}

func intPower(a, b int) int {

	return int(math.Pow(float64(a), float64(b)))

}

func reverseString(s string) string {

	var sNew string
	for i := len(s) - 1; i >= 0; i-- {
		sNew += string(s[i])
	}

	return sNew

}
