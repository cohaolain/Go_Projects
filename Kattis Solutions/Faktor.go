package main

import "fmt"

func main() {

	var a, i int
	fmt.Scanf("%d %d", &a, &i)

	var scientists int
	for true {
		if int(scientists/a)+1 >= i {
			fmt.Println(scientists + 1)
			break
		} else {
			scientists++
		}
	}

}
