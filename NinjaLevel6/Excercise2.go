package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo4(v...)
	y := bar4(v)

	fmt.Println(x)
	fmt.Println(y)
}

func foo4(vals ...int) int {
	sums := 0
	for _, v := range vals {
		sums += v
	}
	return sums
}

func bar4(vals []int) int {
	sums := 0
	for _, v := range vals {
		sums += v
	}
	return sums
}
