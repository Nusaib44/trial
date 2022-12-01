package main

import (
	"fmt"
)

func input(slice []int, err error) []int {
	var value int

	if err != nil {
		return slice
	}

	number, err := fmt.Scanf("%d", &value)
	if number == 1 {
		slice = append(slice, value)
	}

	return input(slice, err)
}

func main() {

	fmt.Println("Enter the values")
	slice := input([]int{}, nil)

	fmt.Println(slice)

}
