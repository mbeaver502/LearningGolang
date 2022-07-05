package main

import "fmt"

func main() {
	y := 1992

	for {
		if y > 2022 {
			break
		}

		fmt.Println(y)
		y++
	}
}
