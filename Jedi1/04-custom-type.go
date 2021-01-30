package main

import "fmt"

type chuchi int

var x chuchi



func main() {
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	x = 42
	fmt.Println(x)
}
