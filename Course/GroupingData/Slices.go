package main

import "fmt"

func main() {
	// x := type{values} ----> composite literal
	x := []int{10, 20, 30, 40} // VALUES OF THE SAME TYPE
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
}
