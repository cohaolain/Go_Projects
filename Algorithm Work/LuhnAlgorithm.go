package main

import "fmt"
import "strconv"

func main() {

	fmt.Print("\n\tWelcome to Ciarán Ó hAoláin's Luhn Algorithm Verification program, written in Go!\n\n\tStart by typing a number, and hitting return!\n\n")

	for {
		var number int
		fmt.Printf("\t")
		fmt.Scanf("%d\n", &number)

		if number != 0 {

			if isValid(number) {
				fmt.Println("\tValid\n\n")
			} else {
				fmt.Println("\tInvalid\n\n")
			}

		}
	}

}

func isValid(num int) bool {

	length := len(strconv.Itoa(num))

	// Create array containing each digit of the number
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = num % 10
		num /= 10
	}

	// Apply the algorithm
	count := 0
	for i := 0; i < len(array); i += 2 {
		count += array[i]
	}
	for i := 1; i < len(array); i += 2 {
		array[i] *= 2
		if array[i] > 9 {
			array[i] -= 9
		}
		count += array[i]
	}

	// Run final check and return the results
	if count%10 == 0 {
		return true
	} else {
		return false
	}

}
