package main

import "fmt"

func main() {
	x := struct {
		first string
		last  string
	}{
		first: "Pablo",
		last:  "Nuñez",
	}

	fmt.Println(x)
}
