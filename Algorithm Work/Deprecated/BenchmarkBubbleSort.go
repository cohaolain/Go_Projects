package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main() {

	/*	Read inputs:
		n = numbers per array
		lim = highest integer in array
		runs = number of arrays to sort
	*/

	var n, lim, runs int
	fmt.Println("\nInput n, lim, runs:\n")
	fmt.Scanf("%d %d %d\n", &n, &lim, &runs)

	var totalTimeBS time.Duration
	var totalTimeBIS time.Duration

	for i := 0; i < runs; i++ {

		// Generate an array of n random integers between 0 and lim
		array := genArray(n, lim)
		array1 := array
		array2 := array

		// Bubble Sort the array, result is discarded
		start := time.Now()
		bubbleSort(array1)
		totalTimeBS += time.Since(start)

		// Use Built-In Sort on the array
		start = time.Now()
		sort.Ints(array2)
		totalTimeBIS += time.Since(start)

		// Print progress
		fmt.Println(float64(i*100)/float64(runs), "%")

	}

	fmt.Println("\nBubble Sort: ", (float64(totalTimeBS.Nanoseconds())/1000000)/float64(runs), "ms")
	fmt.Println("Built-In Sort: ", (float64(totalTimeBIS.Nanoseconds())/1000000)/float64(runs), "ms\n")

}

func bubbleSort(input []int) []int {

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

func genArray(n, lim int) []int {

	rand.Seed(time.Now().UnixNano())
	var array []int

	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(lim))
	}

	return array

}
