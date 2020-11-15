package main

import (
	"fmt"
)

// define the node
type Node struct {
	data int   // Data
	next *Node // Reference to next node
}

type llist struct {
	head *Node
	tail *Node
}

// ADD value to the list
func (l *llist) AddValue(value int) bool {
	node := &Node{data: value}
	if l.head == nil {
		l.head = node
		l.tail = node
		return true
	}
	for l.tail.next != nil {
		l.tail = l.tail.next
	}
	if l.tail.next == nil {
		l.tail.next = node
		return true
	}
	return false
}

func (l *llist) deleteValue(value int) bool {
	prev := l.head
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.data == value {
			l.tail = prev
			l.tail.next = curr.next
			return true
		}
		prev = curr
	}
	return false
}

func (l *llist) insertValue(value, after int) bool {
	node := &Node{data: value}
	temp := l.head
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.data == after {
			temp = curr.next
			curr.next = node
			curr = curr.next
			curr.next = temp
			return true
		}
	}
	return false
}

func (l llist) isEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

// List all values
func (l llist) liatAll() {
	for l.tail = l.head; l.tail != nil; l.tail = l.tail.next {
		fmt.Println(l.tail.data)
	}
}

func main() {
	mylist := &llist{}
	mylist.AddValue(1)
	mylist.AddValue(2)
	mylist.AddValue(3)
	mylist.liatAll()
	mylist.deleteValue(2)
	mylist.liatAll()
	mylist.AddValue(4)
	mylist.liatAll()
	mylist.insertValue(1, 1)
	mylist.liatAll()
	fmt.Println(mylist.isEmpty())
}
