package main

import (
	"fmt"
)

func main() {

	values := make(map[int]int)
	var list []int
	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i<n; i++ {
		var value int
		fmt.Scanf("%d", &value)
		values[value]++
		list = append(list, value)
	}

	record := -1
	for value, count := range values {
		if count < 2 && value > record {
			record = value
		}
	}

	if record != -1 {
		for index, value := range list {
			if value == record {
				fmt.Println(index+1)
				break
			}
		}
	} else {
		fmt.Println("none")
	}

}
