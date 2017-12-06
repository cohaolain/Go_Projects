package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	required := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		count := make(map[string]int)
		for j := 0; j < len(input); j++ {
			if strings.Contains(required, strings.ToLower(input[j:j+1])) {
				count[strings.ToLower(input[j:j+1])]++
			}
		}
		var found []string
		for j := 0; j < len(required); j++ {
			if count[required[j:j+1]] == 0 {
				found = append(found, required[j:j+1])
			}
		}
		if len(found) > 0 {
			fmt.Print("missing ")
			for j := 0; j < len(found); j++ {
				fmt.Print(found[j])
			}
			fmt.Println()
		} else {
			fmt.Println("pangram")
		}
	}
}
