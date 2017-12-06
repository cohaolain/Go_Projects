package main

import (
	"fmt"
	"math"
)

func main() {

	n := inputIntByNewline()

	fmt.Print(math.Sqrt(float64(n))*4)

}

func inputIntByNewline() (n int) {
	fmt.Scanf("%d\n", &n)
	return
}