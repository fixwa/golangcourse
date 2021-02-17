package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	fooo()
	fmt.Println(x)
}

func fooo() {
	fmt.Println("Calling FOOO...")
	x++
}
