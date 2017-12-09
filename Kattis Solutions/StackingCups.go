package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	sort2 "sort"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	sort := make(map[int]string)

	for i:=0; i<n; i++ {

		scanner.Scan()
		var realRadius int
		var colour string
		if radius, err := strconv.Atoi(scanner.Text()); err == nil {
			realRadius = radius/2
			scanner.Scan()
			colour = scanner.Text()
		} else {
			colour = scanner.Text()
			scanner.Scan()
			if radius2, err := strconv.Atoi(scanner.Text()); err == nil {
				realRadius = radius2
			} else {
				// This won't happen
			}
		}
		sort[realRadius]=colour

	}

	array := make([]int, len(sort))

	for i := range sort {
		array = append(array, i)
	}

	sort2.Ints(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(sort[array[i]])
	}

}
