package main

import "fmt"

type chuchi int

var x chuchi
var y int


func main() {
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	x = 42
	fmt.Println(x)


	fmt.Printf("%T \n", y)
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T \n", y)
}
