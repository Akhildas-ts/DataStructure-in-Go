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

func (l *LinkedList) MergeTwoSortArray(list1 *Node, list2 *Node) *Node {

	l1 := list1
	l2 := list2
	var mergedNode *Node
	var prev *Node

	for l1 != nil && l2 != nil {

		var store *Node

		if l1.Data < l2.Data {
			store = l1
			l1 = l1.Next
		} else if l2.Data < l1.Data {
			store = l2
			l2 = l2.Next
		}

		if mergedNode == nil {
			mergedNode = store
			prev = mergedNode
		} else {
			prev.Next = store
			prev = store
		}

	}

	if l1 != nil {
		prev.Next = l1

	} else if l2 != nil {
		prev.Next = l2
	}
	return mergedNode

}

//REMOVE THE LAST N TH ELEMENT OF THE LINKED LIST
// FRIST OF ALL WE ARE FINDING THE LENGHT
// THEN WE ARE GOING TO THE N (target)
//THEN  IF THE TARGET REACH SO THAT TIME WE NEED DELTE THAT NODE

func (l *LinkedList) RemoveLastNelement(head *Node, n int) *Node {
	if head == nil || n <= 0 {
		return head
	}

	// Calculate the length of the linked list
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	// Reset current to the head of the list
	current = head

	// If n is equal to the length of the list, remove the first node
	if n == length {
		head = head.Next
		return head
	}

	// Traverse the list again to find the node before the nth node from the end
	for i := 1; i < length-n; i++ {
		current = current.Next
	}

	// Remove the nth node from the end
	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}

//SWAP THE NODE'S

func (l *LinkedList) SwapPairs(head *Node) *Node {
	l.Head = head

	if head == nil {
		fmt.Println("there is no nodes")
		return nil
	}

	current := head

	var prev *Node
	var newHead *Node

	for current != nil && current.Next != nil {

		firstNode := current
		secondNode := current.Next

		// Adjust the pointers to swap the nodes
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// Update the head of the list if the swap occurs at the beginning
		if current == head {
			head = secondNode
		}

		if prev != nil {
			prev.Next = secondNode
		} else {
			newHead = secondNode
		}
		// Move to the next pair of nodes
		prev = firstNode
		current = firstNode.Next
	}

	return newHead

}

// REMOVE DUPLICATE ELEMENTS

func (l *LinkedList) RemvoeDuplicate(head *Node) {

	store := make(map[int]bool)

	c := head

	if head == nil {
		fmt.Println("there is no more node are there ")
		return
	}

	store[c.Data] = true

	for c.Next != nil {

		if store[c.Next.Data] {
			c.Next = c.Next.Next
		} else {

			store[c.Next.Data] = true
			c = c.Next
		}
	}

	c = head
	l.NodePrint(c)

}

// SORT THE LINKED LIST
func (l *LinkedList) SortList(head *Node) *Node {
	if head == nil {
		// Nothing to sort if the list is empty or has only one node
		return nil
	}

	if head != nil && head == nil {
		return head
	}

	// Iterate through the list using bubble sort algorithm
	for current := head; current != nil; current = current.Next {
		fmt.Println("printing the currentnode", current.Data)
		for nextNode := current.Next; nextNode != nil; nextNode = nextNode.Next {
			fmt.Println("prini the next node", nextNode.Data)
			if current.Data > nextNode.Data {
				// Swap the data of the current node and the next node
				current.Data, nextNode.Data = nextNode.Data, current.Data
			}
		}
	}

	// Print the sorted list
	l.NodePrint(head)
	println()
	return head
}
