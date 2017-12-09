package UsefulFunctions

import (
	"fmt"
	"math"
)

func inputIntByNewline() (n int) {
	fmt.Scanf("%d\n", &n)
	return
}

func inputStringByNewline() (str string) {
	fmt.Scanf("%s\n", &str)
	return
}


func inputArrayIntsByNewline(numberOfInts int) (out []int) {
	for i := 0; i < numberOfInts; i++ {
		out = append(out, inputIntByNewline())
	}
	return
}

func inputArrayStringsByNewline(numberOfStrings int) (out []string) {
	for i := 0; i < numberOfStrings; i++ {
		out = append(out, inputStringByNewline())
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
		return mostBases[:base]
	}
}

func toBase(target string, input int) (output string){
	baseTarget := len(target)
	for remainder := input; true; {
		output += string(target[remainder%baseTarget])
		remainder = (remainder - remainder%baseTarget) / baseTarget
		if remainder == 0 {
			break
		}
	}
	output = reverseString(output)
	return
}

func reverseString(s string) (sNew string) {
	for i := len(s) - 1; i >= 0; i-- {
		sNew += string(s[i])
	}
	return
}
