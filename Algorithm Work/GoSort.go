package main

import (
	"fmt"
	"time"
	"math/rand"
	"sort"
)

func genArray(n, lim int) []int {

	rand.Seed(time.Now().UnixNano())
	var array []int

	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(lim))
	}

	return array

}

func main() {

	array := genArray(500, 50000)

	fmt.Println(array)

	sort.Ints(array)

	fmt.Println(array)

}
