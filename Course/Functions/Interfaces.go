package main

import "fmt"

type person1 struct {
	first string
	last  string
}

type secretAgent1 struct {
	person1
	ltk bool
}

type hotdog int

func (s secretAgent1) speak() {
	fmt.Println("I am", s.first, s.last, " - the Secret Agent speak")
}

func (s person1) speak() {
	fmt.Println("I am", s.first, s.last, " - the Person speak")
}

type human interface {
	speak()
}

func barr(h human) {
	switch h.(type) {
	case person1:
		fmt.Println("PERSON passed into BARR", h.(person1).first)
	case secretAgent1:
		fmt.Println("SecretAgent1 passed into BARR", h.(secretAgent1).first)
	}

	fmt.Println("I was passed into BARR", h)
}

func main() {
	sa1 := secretAgent1{
		person1: person1{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person1{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(p1)

	barr(sa1)
	barr(p1)

	// Conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int = 42
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
