package main

import "fmt"

func main() {

	var runs int
	fmt.Scanf("%d\n", &runs)

	for i := 0; i < runs; i++ {

		var testCase int
		fmt.Scanf("%d\n", &testCase)
		if testCase%2 != 0 {
			fmt.Printf("%d is odd\n", testCase)
		} else {
			fmt.Printf("%d is even\n", testCase)
		}

	}
}
