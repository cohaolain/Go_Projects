package main

import "fmt"

func main() {

	var n, counter int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		var num int
		fmt.Scanf("%d", &num)

		if num < 0 {
			counter++
		}
	}

	fmt.Print(counter)

}
