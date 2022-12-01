package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// ''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''

// Adding value

// __________________________________________________________________________________________________________
// Method1
// *linkedList because to modify [added element go to the linkedList object]
func (a *linkedList) addvalue1(val int) {
	newNode := node{data: val}
	if a.head != nil {
		newNode.next = a.head
		a.head = &newNode
		a.length++
	} else {
		a.head = &newNode
		a.length++
	}
}

// ________________________________________________________________________________________________________
// Method2

func (a *linkedList) add_prepend(b *node) {
	temp := a.head
	a.head = b
	a.head.next = temp
	a.length++

}

// ________________________________________________________________________________________________________
func main() {
	var c int
	myList := linkedList{}
	fmt.Println("Enter a value")
	fmt.Scan(&c)
	myList.addvalue1(c)
	fmt.Println(myList)

	k := &node{data: 4}
	myList.add_prepend(k)
	fmt.Println(myList)
}
