package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		var s1, s2, ans string
		fmt.Scanf("%s\n%s\n", &s1, &s2)

		for j, val := range s1 {

			if uint8(val) == s2[j] {
				ans += "."
			} else {
				ans += "*"
			}

		}
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(ans)
		fmt.Println()

	}

}
