package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	switch {
	case x > 80:
		fmt.Printf("More than 80: %v\n", x)
	case x <= 70 && x >= 50:
		fmt.Printf("Between 70 and 50: %v\n", x)
	default:
		fmt.Printf("Other : %v\n", x)
	}
}
