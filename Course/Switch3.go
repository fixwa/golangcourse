package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny", "bond", "Do No":
		fmt.Println("Miss money or bond or do no")
	case "BOND", "Bond":
		fmt.Println("Bond james")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("1This is the dafault!")
	}
}
