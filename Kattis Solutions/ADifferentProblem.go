package main

import (
	"fmt"
	"math"
)

func main() {

	for {
		var n1, n2 int
		fmt.Scanf("%d %d\n", &n1, &n2)
		if n1 == 0 {
			break
		}
		fmt.Println(int(math.Abs(float64(n1)-float64(n2))))
	}

}
