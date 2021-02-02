package main

import "fmt"

func main() {

	for i := 2020; i >= 1980; i-- {
		fmt.Println(i)
	}

	bd := 1980
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}
