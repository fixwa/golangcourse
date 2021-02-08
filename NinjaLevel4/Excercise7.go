package main

import "fmt"

func main() {
	x := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Hellooooooo, James"},
	}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println("Record: ", i)
		for j, w := range v {
			fmt.Printf("\tIndex position: %v\t Value: %v\n", j, w)
		}
	}
}
