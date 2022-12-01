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

func (a *linkedList) addElement(val1 int) {
	nextNode := node{data: val1}
	add := a.head

	if a.length == 0 {
		a.head = &nextNode
		a.length++
		return
	}

	for i := 0; i < a.length; i++ {

		if add.next == nil {
			add.next = &nextNode
			a.length++
			return
		}
		add = add.next
	}

}
func (a *linkedList) add_prepend(b *node) {
	temp := a.head
	a.head = b
	a.head.next = temp
	a.length++

}

func (l *linkedList) InsertAt(pos int, val2 int) {
	// create a new node
	newNode := node{}
	newNode.data = val2
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.length++
		return
	}
	if pos > l.length {
		return
	}
	// n := l.(pos)
	// newNode.next = n
	// prevNode := l.GetAt(pos - 1)
	// prevNode.next = &newNode
	// l.length++
}

func (a *linkedList) print() {
	ptr := a.head

	for i := 0; i < a.length; i++ {
		fmt.Print("    ", ptr.data)
		ptr = ptr.next

	}
}

func main() {
	myList := linkedList{}

	myList.addElement(33)
	myList.addElement(34)
	myList.addElement(35)

	myList.print()
	myList.print()
}
