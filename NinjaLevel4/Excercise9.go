package main

import "fmt"

func main() {
	var x = map[string][]string{
		"pablo_nunez":       []string{"Age of Empires", "Battlefield 4"},
		"jessica_fernandez": []string{"Cats", "MacBooks"},
		"chavo_ocho":        []string{"Tortas de Jamón", "Comida"},
	}

	fmt.Println(x)

	x["spongebob"] = []string{"Pattys", "Butterflies"}

	for k, v := range x {
		fmt.Println("This is the record for: ", k)

		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}
}
