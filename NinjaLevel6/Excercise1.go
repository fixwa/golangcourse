package main

import "fmt"

func main() {
	fmt.Println(foo3())
	fmt.Println(bar2())
}

func foo3() int {
	return 2
}

func bar2() (int, string) {
	return 99, "ok"
}
