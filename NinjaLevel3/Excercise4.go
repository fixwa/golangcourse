package main

import "fmt"

func main() {
	bd := 1980
	for {
		fmt.Println(bd)
		bd++

		if bd > 2020 {
			break
		}

	}
}
