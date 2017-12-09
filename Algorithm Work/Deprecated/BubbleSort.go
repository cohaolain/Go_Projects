package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	array := inputArray(n)

	fmt.Println("Unsorted:\t", array)

	bubbleSortInts(array)

	fmt.Println("Sorted:\t\t", array)

}

func bubbleSortInts(input []int) []int {

	for cont := true; cont; {
		cont = false
		for i := 0; i < len(input)-1; i++ {
			if input[i+1] < input[i] {
				input[i], input[i+1] = input[i+1], input[i]
				cont = true
			}
		}
	}
	return input

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
