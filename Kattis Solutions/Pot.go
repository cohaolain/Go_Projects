package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	array := inputArray(n)

	answer := 0
	for i := 0; i < n; i++ {
		answer += int(math.Pow(float64(array[i]-array[i]%10)/10, float64(array[i]%10)))
	}

	fmt.Print(answer)

}

func inputArray(n int) []int {

	var array []int
	var input int

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &input)
		array = append(array, input)
	}

	return array

}
