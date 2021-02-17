package main

import "fmt"

func main() {
	x := bars()
	fmt.Printf("%T\n", x)
	fmt.Println(x())
}

func bars() func() int {
	return func() int {
		return 451
	}
}
