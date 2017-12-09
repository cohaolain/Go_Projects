package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)

	var logs [][]string
	var empty []string
	logs = append(logs, empty)

	for count :=0; input.Scan(); {
		line := input.Text()
		if line == "" {
			logs = append(logs, empty)
			count++
			continue
		}
		logs[count] = append(logs[count], line)
	}

	for i := range logs {
		totalCount := 0
		for j := range logs[i] {
			count := 0
			for _, char := range logs[i][j] {
				if string(char) == "*" {
					count++
				}
			}
			for k := 0; k < len(logs[i][j])-count-totalCount; k++ {
				fmt.Print(".")
			}
			for k := 0; k < count; k++ {
				fmt.Print("*")
			}
			for k := 0; k < totalCount; k++ {
				fmt.Print(".")
			}
			totalCount+=count
			fmt.Println()
		}
		fmt.Println()
	}

}
