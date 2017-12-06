package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	for i:=0; i<20; i++ {
		fmt.Printf("%d ", rand.Intn(2))
	}

}
