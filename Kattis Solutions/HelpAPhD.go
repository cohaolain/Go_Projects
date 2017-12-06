		package main

		import (
			"fmt"
			"strconv"
		)

		func main() {

			var n int
			fmt.Scanf("%d\n", &n)
			for i:=0; i<n; i++ {

				var input string
				var pos int
				var skip bool
				fmt.Scanln("%s\n", &input)
				for j:=0; j<len(input); j++ {
					if input[j:j+1]=="+" {
						pos = j
						break
					}
					if j==len(input)-1 {
						skip = true
					}
				}
				if !skip {
					a, _ := strconv.Atoi(input[:pos])
					b, _ := strconv.Atoi(input[pos+1:])
					fmt.Println(a+b)
				} else {
					fmt.Println("skipped")
				}


			}


		}

