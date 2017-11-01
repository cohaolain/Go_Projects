package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		var n2 int
		fmt.Scanf("%d", &n2)
		gnomes := make([]int, n2)

		for j := 0; j < n2; j++ {
			fmt.Scanf("%d", &gnomes[j])
		}

		for j := 0; j < n2; j++ {
			if gnomes[j+1] != gnomes[j]+1 {
				fmt.Println(j + 2)
				break
			}
		}

	}

}
