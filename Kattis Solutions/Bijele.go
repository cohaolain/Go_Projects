package main

import "fmt"

func main() {

	array := inputArray(6)

	fmt.Printf("%d %d %d %d %d %d", 1-array[0], 1-array[1], 2-array[2], 2-array[3], 2-array[4], 8-array[5])

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
