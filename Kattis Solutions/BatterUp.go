package main

import "fmt"

func main() {

	var n int
	var nums []int
	fmt.Scanf("%d\n", &n)

	for i:=0; i<n; i++ {

		var input int
		fmt.Scanf("%d", &input)


		if input != -1 {
			nums = append(nums, input)
		}

	}

	count := 0

	for i := 0; i<len(nums); i++ {

		count += nums[i]

	}

	fmt.Println(float64(count)/float64(len(nums)))


}