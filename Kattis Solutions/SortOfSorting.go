package main

import "fmt"

func main() {

	var n int

	for true {

		fmt.Scanf("%d\n", &n)

		if n == 0 {
			break
		}

		var array []string

		for i := 0; i < n; i++ {

			var s string
			fmt.Scanf("%s\n", &s)
			array = append(array, s)

		}

		bubbleSort(array)

		for i := 0; i < n; i++ {
			fmt.Println(array[i])
		}
		fmt.Println()
	}

}

func bubbleSort(input []string) []string {

	for cont := true; cont; {
		cont = false
		for i := 0; i < len(input)-1; i++ {
			if input[i+1][0:2] < input[i][0:2] {
				input[i], input[i+1] = input[i+1], input[i]
				cont = true
			}
		}
	}
	return input

}
