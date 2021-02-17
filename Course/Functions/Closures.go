package main

import "fmt"

var xx int

func main() {
	fmt.Println(xx)
	xx++
	fmt.Println(xx)

	{
		y := 44
		fmt.Println(y)
	}

	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("-----------")

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++ //has access to the ´x´ above
		return x
	}
}
