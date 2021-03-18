package main

import "fmt"

func main() {
	defer theDefered()
	notDefered()
}

func notDefered() {
	fmt.Println("Not Defered")
}

func theDefered() {
	fmt.Println("The Defered")
}
