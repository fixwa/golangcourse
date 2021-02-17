package main

import "fmt"

func main() {
	foo()
	bar("Pablo")
	s1 := woo("Jessica")
	fmt.Println(s1)

	x, y := tang("Flaming", "Fenix")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from FOO")
}

func bar(str string) {
	fmt.Printf("Hello whatsup: %v\n", str)
}

func woo(str string) string {
	return fmt.Sprintf("Hello whatsup from WOO: %v\n", str)
}

// Return MULTIPLE
func tang(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
