package main

import "fmt"

type Square struct {
}
type Circle struct {
}

func (p Person) speak() {
	fmt.Printf("My name is %s, and my age is %v", p.first, p.age)
}

func main() {
	p := Person{
		first: "Pablo",
		last:  "Nuñez",
		age:   40,
	}
	fmt.Println(p)

	p.speak()
}
