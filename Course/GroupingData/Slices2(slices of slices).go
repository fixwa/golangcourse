package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40}
	fmt.Println(x[1])
	fmt.Println(x)

	// SLICES OF SLICES
	fmt.Println(x[:])
	fmt.Println(x[0:])
	fmt.Println(x[1:])
	fmt.Println(x[2:])
	fmt.Println(x[3:])
	fmt.Println(x[4:])
	//
	fmt.Println(x[1:3])
}
