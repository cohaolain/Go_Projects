package main

import (
	"os"
	"strings"
	"fmt"
	"bufio"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	dict := make(map[string]string)

	for scanner.Scan() {
		inArr := strings.Split(scanner.Text(), " ")
		if inArr[0] == "" {
			break
		}
		dict[inArr[1]]=inArr[0]
	}

	for scanner.Scan() {
		if dict[scanner.Text()] != "" {
			fmt.Println(dict[scanner.Text()])
		} else {
			fmt.Println("eh")
		}
	}

}