package main

import (
	"fmt"
	"sort"
)

func main() {

	array := inputArray(4)

	sort.Ints(array)

	fmt.Println(array[0] * array[2])

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
