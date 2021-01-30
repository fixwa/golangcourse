package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v %v %v \n", x, y, z)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n",y)
	fmt.Printf("%T \n",z)
}
