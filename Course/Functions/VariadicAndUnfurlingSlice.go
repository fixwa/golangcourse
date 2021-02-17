package main

import "fmt"

func main() {
	s := sum(2, 3, 4, 5, 6, 7, 8)
	fmt.Print("The total is: ", s)

	xi := []int{4, 4, 4, 5, 5, 5, 1}
	s2 := sum(xi...)
	fmt.Print("The total is: ", s2)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("For item in index position:", i, "we are now adding:", v, "to the total which is now:", sum)
	}

	return sum
}
