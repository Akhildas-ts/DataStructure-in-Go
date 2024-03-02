package Linklist

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

//TRAVERSE LINKED LIST AND PRINT ELEMENT

func (l LinkedList) TraverseAllNode() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
}

//INSERTION BEGGINING OF THE NODE

func (l *LinkedList) InesrtBeggingOFNode(val int) {
	node1 := &Node{Data: val}

	second := l.Head
	node1.Next = second
	l.Head = node1

}

//INSERTING AT THE END OF THE NODE

func (l *LinkedList) InsertAtEndOfNOde(val int) {
	node1 := &Node{Data: val}

	if l.Head == nil {
		l.Head = node1
		l.Length = 1
		return
	}

	current := l.Head

	for current.Next != nil {

		current = current.Next
	}

	current.Next = node1
	l.Length++

}

//INSERT MIDDLE OF THE NODE

func (l *LinkedList) AddMiddleOfNode(val int) {

	node1 := &Node{Data: val}

	if l.Head == nil {
		l.Head = node1
		return
	}
	current := l.Head

	pos := l.Length / 2

	for i := 1; i < pos; i++ {

		current = current.Next

	}
	node1.Next = current.Next
	current.Next = node1
	current = node1
}

//DELETE NODE

func (l *LinkedList) DeleteNOde(val int) {

	if l.Head.Data == val {
		l.Head = l.Head.Next
		return
	}

	current := l.Head

	for current.Next.Data != val {
		current = current.Next
	}

	current.Next = current.Next.Next

}

//Serach for specific value

func (l *LinkedList) Serachvalue(val int) (int, int, error) {

	if l.Head.Data == val {

		return val, 1, nil
	}

	current := l.Head.Next

	for current.Next != nil {

		if current.Data == val {
			return current.Data, l.Length, nil
		}
		current = current.Next
	}

	return 0, 0, errors.New("there is no value in the node")
}

// CALCULATE THE LENGTH OF THE LIST

func (l *LinkedList) LengthOfList() {

	current := l.Head
	index := 0

	for current != nil {
		index++
		current = current.Next
	}

	fmt.Println("the lenght of the linked List is :", index)
}

//A7 REVERSE THE LIST

//rerver printing the node with recursion 


func (l *LinkedList) ReverserList() {

	current := l.Head

	l.PrintReverse(current)

}

func (l *LinkedList) PrintReverse(n *Node) {

	if n == nil {
		return
	}

	l.PrintReverse(n.Next)
	fmt.Println(n.Data)

}
