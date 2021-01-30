package main

import "fmt"

var x int
var y string
var z bool


func main() {

	fmt.Printf("%v %v %v \n", x, y, z)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n",y)
	fmt.Printf("%T \n",z)

	x = 42
	y = "James Bond"
	z = true

	fmt.Printf("%v %v %v \n", x, y, z)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n",y)
	fmt.Printf("%T \n",z)
}
