package main

import "fmt"

type list struct {
	head *node
	tail *node
	size int
}

type node struct {
	data int
	next *node
}

func (l *list) add(data int) {

	n := &node{data: data}

	if l.head == nil {

		l.head = n
		l.size = 1
	} else {

		l.tail.next = n
		l.size++
	}
	l.tail = n
}

func (l *list) sizeof() int {

	return l.size

}

func (l *list) print() {

	n := l.head

	for n.next != nil {
		fmt.Printf("%v ->", n.data)
		n = n.next
	}
	fmt.Printf("%v", n.data)
	fmt.Printf("-> nil")

}

func main() {

	l := &list{}
	l.add(1)
	l.add(2)
	l.print()
	k := l.sizeof()

	fmt.Println(" ")
	fmt.Printf("size of list: %v", k)

}
