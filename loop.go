package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 3; i++ {

		for j := 0; j < i*2; j++ {
			fmt.Printf("* ")
		}
		println("")

		for k := 0; k < i*3; k++ {
			if i == 3 {
				break
			} else {
				fmt.Printf("* \n")
			}
		}
	}

}
