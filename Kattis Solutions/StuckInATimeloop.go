package main

import "fmt"

func main() {

	var input int
	fmt.Scanf("%d", &input)

	for i := 0; i < input; i++ {
		fmt.Printf("%d Abracadabra\n", i+1)
	}

}
