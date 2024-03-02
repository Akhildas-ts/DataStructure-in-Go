package Linklist

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

//TRAVERSE LINKED LIST AND PRINT ELEMENT

func (l LinkedList) TraverseAllNode() {
	current := l.head

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}

//INSERTION BEGGINING OF THE NODE

func (l *LinkedList) InesrtBeggingOFNode(val int) {
	node1 := &Node{data: val}

	second := l.head
	node1.next = second
	l.head = node1

}

//INSERTING AT THE END OF THE NODE

func (l *LinkedList) InsertAtEndOfNOde(val int) {
	node1 := &Node{data: val}

	if l.head == nil {
		l.head = node1
		l.length = 1
		return
	}

	current := l.head

	for current.next != nil {

		current = current.next
	}

	current.next = node1
	l.length++

}

//INSERT MIDDLE OF THE NODE

func (l *LinkedList) AddMiddleOfNode(val int) {

	node1 := &Node{data: val}

	if l.head == nil {
		l.head = node1
		return
	}
	current := l.head

	pos := l.length / 2

	for i := 1; i < pos; i++ {

		current = current.next

	}
	node1.next = current.next
	current.next = node1
	current = node1
}

//DELETE NODE

func (l *LinkedList) DeleteNOde(val int) {

	if l.head.data == val {
		l.head = l.head.next
		return
	}

	current := l.head

	for current.next.data != val {
		current = current.next
	}

	current.next = current.next.next

}

//Serach for specific value

func (l *LinkedList) Serachvalue(val int) (int, int, error) {

	if l.head.data == val {

		return val, 1, nil
	}

	current := l.head

	for current.next != nil {

		if current.data == val {
			return current.data, l.length, nil
		}
		current = current.next
	}

	return 0, 0, errors.New("there is no value in the node")
}

// CALCULATE THE LENGTH OF THE LIST

func (l *LinkedList) LengthOfList() {

	current := l.head
	index := 0

	for current != nil {
		index++
		current = current.next
	}

	fmt.Println("the lenght of the linked List is :", index)
}

//A7 REVERSE THE LIST

func (l *LinkedList) ReverserList() {

	current := l.head

	l.PrintReverse(current)

}

func (l *LinkedList) PrintReverse(n *Node) {

	if n == nil {
		return
	}

	l.PrintReverse(n.next)
	fmt.Println(n.data)

}
