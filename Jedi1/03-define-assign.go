package main

import "fmt"

var x int = 42
var y string = "James bond"
var z bool = true


func main() {

	fmt.Printf("%v %v %v \n", x, y, z)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n",y)
	fmt.Printf("%T \n",z)

}
