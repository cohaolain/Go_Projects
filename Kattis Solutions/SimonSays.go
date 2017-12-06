package main

import ( "fmt"
		"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Scanf("%d\n", &n)
	for i:=0; i<n; i++ {

		scanner.Scan()
		line := scanner.Text()
		if len(line)<10 {
			continue
		}
		if line[:10]=="Simon says" {
			fmt.Println(line[11:])
		}

	}


}


