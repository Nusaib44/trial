package main

import "fmt"

type stack struct {
	data []int
}

// push
func (a *stack) push(val1 int) {

	a.data = append(a.data, val1)
	// fmt.Println(a.data)
}

// pop
func (a *stack) pop() {

	last := len(a.data) - 1
	// ele := a.data[last]
	a.data = a.data[:last]
	fmt.Println(a.data)
}

// func inputslice(x []int, err error) []int {
// 	if err != nil {
// 		return x
// 	}
// 	var d int
// 	n, err := fmt.Scanf("%d", &d)
// 	if n == 1 {
// 		x = append(x, d)
// 	}
// 	return inputslice(x, err)
// }

// func option() {
// 	var a int
// 	mystack := stack{}

// 	fmt.Println("choose a option ")
// 	fmt.Println("1 for deletion ")
// 	fmt.Println("2 for enter other values")
// 	fmt.Scanln(&a)

// 	if a == 1 {
// 		mystack.pop()
// 	}
// 	if a == 2 {
// 		x := inputslice([]int{}, nil)
// 		mystack.push(x...)
// 	}

// }

func (s *stack) print() {
	v := s.data
	fmt.Println(v)

}

func main() {
	mystack := stack{}

	fmt.Println("Hello welcome....")
	// x := inputslice([]int{}, nil)
	// mystack.push(x...)
	// option()
	mystack.push(2)
	mystack.push(4)
	mystack.push(6)
	mystack.push(8)
	mystack.print()
	mystack.pop()
}
