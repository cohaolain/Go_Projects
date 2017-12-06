package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	for i:=0; i<n; i++ {
		var a,b,c int
		fmt.Scanf("%d %d %d\n", &a, &b, &c)
		if a+b==c || a-b==c || b-a == c || a*b==c || float32(b)/float32(a)==float32(c) || float32(a)/float32(b)==float32(c) {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}

}
