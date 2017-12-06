package main

import "fmt"

func main() {

	var hour, minute int
	fmt.Scanf("%d %d", &hour, &minute)

	if minute >=45 {
		minute-=45
	} else if hour > 0 {
		hour--
		minute += 15
	} else {
		hour = 23
		minute += 15
	}

	fmt.Println(hour, minute)

}
