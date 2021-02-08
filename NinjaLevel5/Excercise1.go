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

	for i, val := range x.favoriteIceCreamFlavors {
		fmt.Println(i, val)
	}

	fmt.Println("--------------")

	for i, val := range y.favoriteIceCreamFlavors {
		fmt.Println(i, val)
	}

}
