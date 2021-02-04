package main

import "fmt"

func main() {
	x := map[string]int{
		"pablo":   100,
		"jessica": 200,
		"hibari":  300,
	}

	fmt.Println(x)

	x["wachin"] = 400

	if v, ok := x["wachin"]; ok {
		fmt.Println(x["wachin"])
		fmt.Println(v)
	}

	for k, v := range x {
		fmt.Println(k, v)
	}

	xi := []int{40, 41, 42, 43, 44, 45}

	for k, v := range xi {
		fmt.Println(k, v)
	}
}
