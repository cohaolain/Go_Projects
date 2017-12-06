package main

import "fmt"

func inputArrayInts(numberOfInts int) (out []int) {
	for i := 0; i < numberOfInts; i++ {
		out = append(out, inputInt())
	}
	return
}

func inputIntByNewline() (n int) {
	fmt.Scanf("%d\n", &n)
	return
}

func inputInt() (n int) {
	fmt.Scanf("%d", &n)
	return
}

func arrayAverage(in []int) (mean float64) {

	var total float64
	for _, val := range in {
		total += float64(val)
	}
	mean = total/float64(len(in))
	return

}

func main() {

	n:=inputIntByNewline()
	var array [][]int

	for i:=0; i<n; i++ {
		elements := inputInt()
		array = append(array, inputArrayInts(elements))
	}

	for _, val := range array {
		average := arrayAverage(val)
		var count float64
		for _, val2 := range val {
			if float64(val2)>average {
				count++
			}
		}
		fmt.Printf("%.3f%%\n", (count*100)/float64(len(val)))
	}

}

