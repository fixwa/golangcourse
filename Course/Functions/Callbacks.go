package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sums(ii...)
	fmt.Println("Sum of all the numbers: ", x)

	s2 := sumEven(sums, ii...)
	fmt.Println("Sum of EVEN numebrs: ", s2)

	s3 := sumOdds(sums, ii...)
	fmt.Println("Sum of ODD numebrs: ", s3)
}

func sums(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sumEven(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func sumOdds(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
