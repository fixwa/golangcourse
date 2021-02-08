package main

import "fmt"

type normalPerson struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	normalPerson
	licenseToKill bool
}

func main() {

	p1 := secretAgent{
		normalPerson: normalPerson{
			first: "James",
			last:  "Bond",
			age:   22,
		},
		licenseToKill: true,
	}

	p2 := secretAgent{
		normalPerson: normalPerson{
			first: "Miss",
			last:  "Moneypenny",
			age:   27,
		},
		licenseToKill: false,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.licenseToKill)
	fmt.Println(p2.last, p2.licenseToKill)
}
