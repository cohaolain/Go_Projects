package main

import "fmt"

func main() {

	var x, y, n int
	fmt.Scanf("%d %d %d", &x, &y, &n)

	for i := 1; i <= n; i++ {

		if i%x == 0 && i%y == 0 {
			fmt.Println("FizzBuzz")
		} else if i%x == 0 {
			fmt.Println("Fizz")
		} else if i%y == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

}
