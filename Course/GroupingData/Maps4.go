package main

import "fmt"

func main() {
	x := map[string]int{
		"pablo":   100,
		"jessica": 200,
		"hibari":  300,
	}

	fmt.Println(x)

	delete(x, "pablo")

	fmt.Println(x)
}
