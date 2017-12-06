// ClassRep.go
// Monte Carlo it, because irony. Best of ten million, anyone?

package Side_Projects

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	names := []string{"Name 1", "Name 2", "Name 3", "Name 4", "Name 5", "Name 6"}
	n := 10000000 // Ten million times is probably enough

	count := make(map[string]int)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		rand := rand.Intn(6)
		count[names[rand]]++
	}

	var scores []int
	for _, score := range count {
		scores = append(scores, score)
	}
	sort.Ints(scores)

	if scores[len(scores)-1] == scores[len(scores)-2] { // Check if the highest and second highest scores are equal
		fmt.Println("There's been a tie!") // This is extremely unlikely
	} else {
		for name, score := range count {
			if scores[len(scores)-1] == score {
				fmt.Println(name)
			}
		}
	}

}