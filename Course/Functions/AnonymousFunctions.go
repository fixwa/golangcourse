package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function was called ...")

	}()

	func(x int) {
		fmt.Println("Anonymous function was called with an arg: ", x)

	}(4299)
}
