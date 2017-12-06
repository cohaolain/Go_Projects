package main

import "fmt"

func main() {

	cups := []bool{true, false, false}

	var moves string
	fmt.Scanf("%s", &moves)

	for i := range moves {

		switch moves[i:i+1] {
		case "A":
			cups[0], cups[1] = cups[1], cups[0]
		case "B":
			cups[1], cups[2] = cups[2], cups[1]
		case "C":
			cups[0], cups[2] = cups[2], cups[0]
		}

	}

	for i := range cups {

		if cups[i] {
			fmt.Println(i+1)
			break
		}

	}

}
