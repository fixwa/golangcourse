package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40}
	y := []int{345, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// DELETE!
	x = append(x[:3], x[7:]...)
	fmt.Println(x)
}
