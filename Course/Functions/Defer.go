package main

import "fmt"

func main() {
	defer def1()
	def2()
}

func def1() {
	fmt.Println("Calling Def1")
}

func def2() {
	fmt.Println("Calling Def2")
}
