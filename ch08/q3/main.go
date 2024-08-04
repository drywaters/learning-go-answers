package main

import "fmt"

type LinkedList[T comparable] struct {
	Root *Node[T]
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// Add appends a node at the end of the linked list
func (l *LinkedList[T]) Add(v T) {
	newNode := Node[T]{
		Value: v,
	}

	if l.Root == nil {
		l.Root = &newNode
		return
	}

	n := l.Root
	for {
		if n.Next == nil {
			n.Next = &newNode
			return
		}
		n = n.Next
	}
}

// Insert inserts a node at the given index
// Insert will panic if index location does not exist
// in the LinkedList
func (l *LinkedList[T]) Insert(v T, index int) {
	if index < 0 {
		panic("insert is out of bounds")
	}

	newNode := Node[T]{
		Value: v,
	}
	if index == 0 {
		newNode.Next = l.Root
		l.Root = &newNode
		return
	}

	n := l.Root
	for i := 0; i < index-1; i++ {
		n = n.Next
		if n == nil {
			panic("insert is out of bounds")
		}
	}

	oldNext := n.Next
	n.Next = &newNode
	newNode.Next = oldNext
}

// Index finds the first index location of a value
func (l *LinkedList[T]) Index(v T) int {
	index := 0
	n := l.Root
	for {
		if n == nil {
			return -1
		} else if n.Value == v {
			return index
		}
		index++
		n = n.Next
	}
}

func main() {
	ln := LinkedList[int]{}
	ln.Add(10)
	ln.Add(20)
	ln.Add(50)
	ln.Insert(30, 0)
	ln.Insert(40, 2)
	// 30 10 40 20 50

	fmt.Println("10 found at", ln.Index(10))
	fmt.Println("20 found at", ln.Index(20))
	fmt.Println("30 found at", ln.Index(30))
	fmt.Println("40 found at", ln.Index(40))
	fmt.Println("50 found at", ln.Index(50))

	ls := LinkedList[string]{}
	ls.Add("Bob")
	ls.Add("Builder")
	ls.Add("Bobby")
	ls.Insert("Other", 0)
	ls.Insert("Name", 0)

	fmt.Println("Name found at", ls.Index("Name"))
	fmt.Println("Other found at", ls.Index("Other"))
	fmt.Println("Bob found at", ls.Index("Bob"))
	fmt.Println("Builder found at", ls.Index("Builder"))
	fmt.Println("Bobby found at", ls.Index("Bobby"))
}
