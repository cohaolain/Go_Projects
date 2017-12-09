package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	var parts []string
	for input.Scan() {
		parts = append(parts, input.Text())
	}

	answers := make(map[string]int)

	for _, val1 := range parts {
		for _, val2 := range parts {
			if val1 != val2 {
				answers[val1+val2]++
			}
		}
	}

	var sortAns []string
	for ans := range answers {
		sortAns = append(sortAns, ans)
	}

	sort.Strings(sortAns)

	for _, val := range sortAns {
		fmt.Println(val)
	}

}
