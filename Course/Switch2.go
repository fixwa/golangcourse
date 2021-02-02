package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print 2")
	case (3 == 3):
		fmt.Println("Prints")
		fallthrough
	case (4 == 4):
		fmt.Println("Also TRUE, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("Not true 1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")

	}

	switch {
	case false:
		fmt.Println("Not print")
	case 1 == 2:
		fmt.Println("1 equals 2?")
	default:
		fmt.Println("1This is the dafault!")
	}
}
