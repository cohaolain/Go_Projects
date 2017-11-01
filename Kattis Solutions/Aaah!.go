package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1, s2 string
	fmt.Scanf("%s\n%s\n", &s1, &s2)
	c1 := strings.Count(s1, "a")
	c2 := strings.Count(s2, "a")

	if c1 >= c2 {

		fmt.Print("go")

	} else {

		fmt.Println("no")

	}

}
