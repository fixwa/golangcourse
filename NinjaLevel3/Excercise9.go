package main

import (
	"fmt"
)

func main() {
	favSport := "paraglide"

	switch favSport {
	case "paraglide":
		fmt.Println("So you like PARAGLIDE... I see...")
	case "football":
		fmt.Println("I can see that you like Football. Ok.")
	default:
		fmt.Printf("I don't know about that sport: %v\n", favSport)
	}
}
