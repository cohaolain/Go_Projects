package main

import "fmt"

func main() {

	var input string
	fmt.Scanf("%s", &input)
	count := make(map[string]int)

	for _, char := range input {
		count[string(char)]++
	}

	remove := 0
	for i := range count {
		if count[i]%2!=0 {
			remove++
		}
	}

	if remove != 0 {
		fmt.Print(remove-1)
	} else {
		fmt.Print(0)
	}

}