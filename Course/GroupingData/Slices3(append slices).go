package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{345, 456, 678, 987}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
}
