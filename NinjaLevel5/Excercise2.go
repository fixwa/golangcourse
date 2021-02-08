package main

import "fmt"

type person struct {
	first                   string
	last                    string
	favoriteIceCreamFlavors []string
}

func main() {

	x := person{
		first:                   "Pablo",
		last:                    "Nuñez",
		favoriteIceCreamFlavors: []string{"Cherry", "Orange", "Strawberry"},
	}

	y := person{
		first:                   "Jessica",
		last:                    "Meriadoc",
		favoriteIceCreamFlavors: []string{"Chocolate", "Cream", "Oreo"},
	}

	m := map[string]person{
		x.first: x,
		y.first: y,
	}

	for _, val := range m {
		fmt.Println(val.first)
		fmt.Println(val.last)
		for i, vv := range val.favoriteIceCreamFlavors {
			fmt.Println(i, vv)
		}

		fmt.Println("------")
	}

}
