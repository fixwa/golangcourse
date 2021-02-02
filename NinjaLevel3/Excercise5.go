package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		mod := i % 4
		fmt.Printf("For value %v, the modulo is %v\n", i, mod)
	}
}
