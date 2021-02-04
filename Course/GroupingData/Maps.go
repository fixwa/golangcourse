package main

import "fmt"

func main() {
	x := map[string]int{
		"pablo":   100,
		"jessica": 200,
		"hibari":  300,
	}

	fmt.Println(x)
	fmt.Println(x["jessica"])

	// not existing key
	fmt.Println(x["WHAT"])

	v, ok := x["WHAT"]
	fmt.Println(v, ok)

	if v, ok := x["WHATS"]; ok {
		fmt.Println("WHATS exists!", v)
	}

}
