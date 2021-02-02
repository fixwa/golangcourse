package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("Value was 40")
	} else if x == 41 {
		fmt.Println("Value was 41")
	} else {
		fmt.Println("Value not 40")
	}
}
