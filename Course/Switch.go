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
	}

}
