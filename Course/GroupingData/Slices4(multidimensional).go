package main

import "fmt"

func main() {
	x := []string{"Pablo", "Nuñez", "Apple"}
	y := []string{"Jessica", "Fern", "Kiwi"}
	fmt.Println(x)
	fmt.Println(y)
	z := [][]string{x, y}
	fmt.Println(z)

}
