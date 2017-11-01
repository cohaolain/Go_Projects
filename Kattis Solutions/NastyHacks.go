package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		var r, e, c int
		fmt.Scanf("%d %d %d", &r, &e, &c)

		if e-c > r {
			fmt.Println("advertise")
		} else if e-c == r {
			fmt.Println("does not matter")
		} else {
			fmt.Println("do not advertise")
		}

	}

}
