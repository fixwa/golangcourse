package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	if x > 50 {
		fmt.Printf("Value is bigger than 50: [%v]\n", x)
	} else if x >= 30 && x <= 40 {
		fmt.Printf("Value is between 30 and 40: [%v]\n", x)
	} else {
		fmt.Printf("Value is smaller than 50: [%v]\n", x)
	}
}
